package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7b80liq2w2ix0p4")
		if err != nil {
			return err
		}

		// update
		edit_thumbnail := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1saqexoz",
			"name": "thumbnail",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"image/jpeg",
					"image/png"
				],
				"thumbs": [
					"250x250f"
				],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_thumbnail); err != nil {
			return err
		}
		collection.Schema.AddField(edit_thumbnail)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7b80liq2w2ix0p4")
		if err != nil {
			return err
		}

		// update
		edit_thumbnail := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1saqexoz",
			"name": "thumbnail",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"image/jpeg",
					"image/png"
				],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_thumbnail); err != nil {
			return err
		}
		collection.Schema.AddField(edit_thumbnail)

		return dao.SaveCollection(collection)
	})
}
