package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"labl/frontend"
	_ "labl/migrations"
	"labl/pkg/render"
	"labl/pkg/search"
	"labl/pkg/templates"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

func SPAMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				if he.Code == http.StatusNotFound {
					c.Request().URL.Path = "/"
					return next(c)
				}
			}
		}
		return err
	}
}

type RenderRequest struct {
	ID       string  `json:"id"`
	Gap      float64 `json:"gap"`
	Margin   float64 `json:"margin"`
	Quantity int     `json:"quantity"`
	Size     struct {
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	} `json:"size"`
	Inputs  render.Inputs `json:"inputs"`
	Outline bool          `json:"outline"`
}

type ImagesRequest struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
	URL  string `json:"url"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func addImages(dao *daos.Dao, appFS *filesystem.System, logger *slog.Logger) error {
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		return nil
	}

	err := filepath.WalkDir("images", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if ext != ".jpg" && ext != ".png" {
			logger.Info("Invalid image file extension, skipping", "file", path)
			return nil
		}
		fileName := d.Name()
		name := strings.TrimSuffix(fileName, ext)
		tag := filepath.Base(filepath.Dir(path))

		foundRecord, _ := dao.FindFirstRecordByFilter("images", "name = {:name} && tag = {:tag}", dbx.Params{"name": name, "tag": tag})

		collection, err := dao.FindCollectionByNameOrId("images")
		if err != nil {
			return err
		}
		record := models.NewRecord(collection)
		if foundRecord != nil {
			record.MarkAsNotNew()
			record.SetId(foundRecord.Id)

			// delete the old file
			oldFileKey := foundRecord.BaseFilesPath() + "/" + foundRecord.GetString("image")
			if err := appFS.Delete(oldFileKey); err != nil {
				return err
			}
		} else {
			record.RefreshId()
		}
		record.Set("name", name)
		record.Set("tag", tag)

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		// generate file metadata
		recordFile, err := filesystem.NewFileFromBytes(data, name+ext)
		if err != nil {
			return err
		}

		record.Set("image", recordFile.Name)

		// upload the file
		fileKey := record.BaseFilesPath() + "/" + recordFile.Name

		if err := appFS.UploadFile(recordFile, fileKey); err != nil {
			return err
		}

		if err := dao.SaveRecord(record); err != nil {
			appFS.Delete(fileKey)
			return err
		}

		return nil
	})

	return err
}

func addTemplates(dao *daos.Dao, fs *filesystem.System, logger *slog.Logger) error {
	templatesDir, err := os.ReadDir("templates")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		logger.Error("Failed to read templates directory", "error", err)
		return err
	}

	for _, file := range templatesDir {
		if file.IsDir() {
			continue
		}
		var name, ext string
		fileName := file.Name()
		ext = filepath.Ext(fileName)
		name = strings.TrimSuffix(fileName, ext)
		if ext != ".json" {
			continue
		}

		foundRecord, _ := dao.FindFirstRecordByFilter("templates", "name = {:name}", dbx.Params{"name": name})

		var template templates.Template
		file, err := os.Open("templates/" + fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		err = json.NewDecoder(file).Decode(&template)
		if err != nil {
			logger.Error("Failed to decode template file", "error", err)
			continue
		}
		encoded, err := json.Marshal(template)
		if err != nil {
			return err
		}

		collection, err := dao.FindCollectionByNameOrId("templates")
		if err != nil {
			return err
		}
		record := models.NewRecord(collection)
		if foundRecord != nil {
			record.MarkAsNotNew()
			record.SetId(foundRecord.Id)

			if foundRecord.GetString("thumbnail") != "" {
				// delete the old thumbnail
				oldThumbnailKey := foundRecord.BaseFilesPath() + "/" + foundRecord.GetString("thumbnail")
				if err := fs.Delete(oldThumbnailKey); err != nil {
					return err
				}
			}
		} else {
			record.RefreshId()
		}
		record.Set("name", name)
		record.Set("data", encoded)

		// check if file with same name but jpg or png extension exists
		name = strings.TrimSuffix(fileName, ".json")
		var thumbnailPath string
		if _, err := os.Stat("templates/" + name + ".jpg"); err == nil {
			thumbnailPath = name + ".jpg"
		} else if _, err := os.Stat("templates/" + name + ".png"); err == nil {
			thumbnailPath = name + ".png"
		}

		if thumbnailPath != "" {
			thumbnail, err := os.Open("templates/" + thumbnailPath)
			if err != nil {
				logger.Error("Failed to open thumbnail file", "error", err)
				return err
			}
			defer thumbnail.Close()

			thumbnailData, err := io.ReadAll(thumbnail)
			if err != nil {
				logger.Error("Failed to read thumbnail file", "error", err)
				return err
			}

			thumbnailFile, err := filesystem.NewFileFromBytes(thumbnailData, filepath.Base(thumbnailPath))
			if err != nil {
				return err
			}
			key := record.BaseFilesPath() + "/" + thumbnailFile.Name
			if err := fs.UploadFile(thumbnailFile, key); err != nil {
				return err
			}

			record.Set("thumbnail", thumbnailFile.Name)
		} else {
			logger.Warn("Thumbnail not found for template", "name", name)
		}

		if err := dao.SaveRecord(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load .env file", "error", err)
	}

	app := pocketbase.New()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	searchEnabled := false
	googleSearchKey := os.Getenv("GOOGLE_SEARCH_KEY")
	googleSearchEngineID := os.Getenv("GOOGLE_SEARCH_ENGINE_ID")
	if googleSearchKey == "" || googleSearchEngineID == "" {
		slog.Error("image search credentials not found, search functionality will be disabled")
	} else {
		slog.Info("image search credentials found, search functionality will be enabled", "key", googleSearchKey, "engine_id", googleSearchEngineID)
		searchEnabled = true
	}

	searcher := search.NewSearcher(googleSearchKey, googleSearchEngineID)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		router := e.Router
		dao := app.Dao()
		fs, err := app.NewFilesystem()
		if err != nil {
			return err
		}
		logger := app.Logger()

		err = addImages(dao, fs, logger)
		if err != nil {
			return err
		}
		err = addTemplates(dao, fs, logger)
		if err != nil {
			return err
		}

		e.Router.Pre(SPAMiddleware)
		e.Router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				return strings.HasPrefix(c.Request().URL.Path, "/_/")
			},
		}))

		subFS := echo.MustSubFS(frontend.Assets, "build")
		e.Router.StaticFS("/", subFS)

		router.POST("/render", func(c echo.Context) error {
			var req RenderRequest
			if err := c.Bind(&req); err != nil {
				return err
			}

			templateRecord, err := dao.FindRecordById("templates", req.ID)
			if err != nil {
				return c.JSON(404, ErrorResponse{404, "Template not found", nil})
			}

			templateData := templateRecord.GetString("data")
			var template templates.Template
			if err := json.Unmarshal([]byte(templateData), &template); err != nil {
				logger.Error("Failed to unmarshal template", "error", err)
				return c.JSON(500, ErrorResponse{500, "Failed to unmarshal template", err})
			}

			for name, resource := range template.Resources {
				if resource.Source.Type == templates.SourceTypeInput {
					input, ok := req.Inputs[name]
					if !ok {
						return c.JSON(400, ErrorResponse{400, "Input not found", nil})
					}
					if resource.Type == templates.ResourceTypeImage {
						inputString, ok := input.(string)
						if !ok {
							return c.JSON(400, ErrorResponse{400, "Image input is not a string", nil})
						}

						image, err := dao.FindRecordById("images", inputString)
						if err != nil {
							return c.JSON(404, ErrorResponse{404, "Image not found", nil})
						}

						key := image.BaseFilesPath() + "/" + image.GetString("image")
						blob, err := fs.GetFile(key)
						if err != nil {
							return c.JSON(500, ErrorResponse{500, "Failed to get image", err})
						}
						defer blob.Close()

						data, err := io.ReadAll(blob)
						if err != nil {
							return c.JSON(500, ErrorResponse{500, "Failed to read image", err})
						}
						req.Inputs[name] = data
					}
				}
			}

			pdf := fpdf.New("P", "mm", "A4", "")
			renderer := render.TemplateRenderer{
				Template:  &template,
				PDF:       pdf,
				Inputs:    req.Inputs,
				Resources: render.Resources{},
				Gap:       req.Gap,
				Margin:    req.Margin,
				Quantity:  req.Quantity,
				Size: render.Size{
					Width:  req.Size.Width,
					Height: req.Size.Height,
				},
				Outline: req.Outline,
			}

			err = renderer.Render()
			if err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to render template", err})
			}

			var pdfData bytes.Buffer
			err = pdf.Output(&pdfData)
			if err != nil {
				logger.Error("Failed to output PDF", "error", err)
				return c.JSON(500, ErrorResponse{500, "Failed to output PDF", err})
			}

			return c.Blob(200, "application/pdf", pdfData.Bytes())
		})

		router.GET("/search", func(c echo.Context) error {
			if !searchEnabled {
				return c.JSON(400, ErrorResponse{400, "Search functionality is disabled", nil})
			}

			query := c.QueryParam("q")
			if query == "" {
				return c.JSON(400, ErrorResponse{400, "Query parameter is required", nil})
			}
			transparent := c.QueryParam("transparent") == "true"
			start, err := strconv.Atoi(c.QueryParam("start"))
			if err != nil {
				start = 1
			}

			var resp *search.SearchResponse

			resp, err = searcher.SearchImages(query, start, transparent)
			if err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to search images", err})
			}

			return c.JSON(200, resp)
		})

		router.POST("/images", func(c echo.Context) error {
			var req ImagesRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(400, ErrorResponse{400, "Failed to bind request", err})
			}

			if req.Name == "" || req.Tag == "" || req.URL == "" {
				return c.JSON(400, ErrorResponse{400, "Name, tag and URL are required", nil})
			}

			collection, err := dao.FindCollectionByNameOrId("images")
			if err != nil {
				return c.JSON(500, ErrorResponse{500, "internal error", err})
			}
			record := models.NewRecord(collection)
			record.RefreshId()
			record.Set("name", req.Name)
			record.Set("tag", req.Tag)

			resp, err := http.Get(req.URL)
			if err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to fetch image", err})
			}
			defer resp.Body.Close()

			ext := filepath.Ext(req.URL)

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to read image data", err})
			}
			file, err := filesystem.NewFileFromBytes(data, req.Name+ext)
			if err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to create file", err})
			}
			key := record.BaseFilesPath() + "/" + file.Name
			if err := fs.UploadFile(file, key); err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to upload file", err})
			}

			record.Set("image", file.Name)

			if err := dao.SaveRecord(record); err != nil {
				return c.JSON(500, ErrorResponse{500, "Failed to save record", err})
			}

			return c.JSON(200, record)
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
