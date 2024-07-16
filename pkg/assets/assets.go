package assets

import "embed"

var Assets = map[string][]byte{}

//go:embed assets/*
var fs embed.FS

func init() {
	files, err := fs.ReadDir("assets")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		data, err := fs.ReadFile("assets/" + file.Name())
		if err != nil {
			panic(err)
		}

		Assets[file.Name()] = data
	}
}
