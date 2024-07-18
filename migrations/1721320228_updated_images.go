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

		collection, err := dao.FindCollectionByNameOrId("jusg9jbfpx83kp9")
		if err != nil {
			return err
		}

		// update
		edit_image := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ghhkkn80",
			"name": "image",
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
		}`), edit_image); err != nil {
			return err
		}
		collection.Schema.AddField(edit_image)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("jusg9jbfpx83kp9")
		if err != nil {
			return err
		}

		// update
		edit_image := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ghhkkn80",
			"name": "image",
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
					"150x150f"
				],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_image); err != nil {
			return err
		}
		collection.Schema.AddField(edit_image)

		return dao.SaveCollection(collection)
	})
}
