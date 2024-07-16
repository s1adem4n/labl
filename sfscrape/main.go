package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const BaseURL = "https://www.sweet-family.de/"

type Design struct {
	Source string
	Title  string
}

type Category struct {
	Title     string
	Thumbnail string
	Designs   []Design
}

func fixName(name string) string {
	ext := filepath.Ext(name)
	name = strings.TrimSuffix(name, ext)

	replacer := strings.NewReplacer(
		"ue", "ü",
		"ae", "ä",
		"oe", "ö",
		"_", " ",
	)
	name = replacer.Replace(name)

	// now remove all numbers
	replacer = strings.NewReplacer(
		"0", "",
		"1", "",
		"2", "",
		"3", "",
		"4", "",
		"5", "",
		"6", "",
		"7", "",
		"8", "",
		"9", "",
	)
	name = replacer.Replace(name)
	name = strings.TrimSpace(name)

	return name
}

func fixNames(names []string) []string {
	nameCount := make(map[string]int)
	fixedNames := make([]string, len(names))

	for i, name := range names {
		name = fixName(name)

		count := nameCount[name]
		if count > 0 {
			fixedNames[i] = fmt.Sprintf("%s (%d).png", name, count)
		} else {
			fixedNames[i] = name + ".png"
		}
		nameCount[name]++
	}

	return fixedNames
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		slog.Error("can't open data file", "err", err)
		return
	}
	defer file.Close()

	var categories []Category
	err = json.NewDecoder(file).Decode(&categories)
	if err != nil {
		slog.Error("can't decode data file", "err", err)
		return
	}

	ticker := time.NewTicker(200 * time.Millisecond)

	for _, category := range categories {
		<-ticker.C

		title := fixName(category.Title)
		err := os.MkdirAll("output/"+title, 0755)
		if err != nil {
			slog.Error("can't create category directory", "err", err)
			return
		}

		var names []string
		for _, design := range category.Designs {
			names = append(names, design.Title)
		}
		fixedNames := fixNames(names)

		for i, design := range category.Designs {
			res, err := http.Get(BaseURL + design.Source)
			if err != nil {
				slog.Error("can't get design", "err", err)
				return
			}
			defer res.Body.Close()

			file, err := os.Create(
				filepath.Join("output", title, fixedNames[i]),
			)
			if err != nil {
				slog.Error("can't create design file", "err", err)
				return
			}
			defer file.Close()

			_, err = file.ReadFrom(res.Body)
			if err != nil {
				slog.Error("can't save design file", "err", err)
				return
			}

			slog.Info("saved design", "category", title, "design", fixedNames[i])
		}
	}
}
