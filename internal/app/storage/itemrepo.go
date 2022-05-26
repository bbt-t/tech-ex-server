package storage

import (
	"log"
	"tech-ex-server/internal/app/model"
)

type ItemRepository struct {
	storage *Storage
}

func (i *ItemRepository) CreateItem(modelObj model.Item) error {
	defer i.storage.Close()
	if err := i.storage.Open(); err != nil {
		log.Fatal(err)
	}

	err := i.storage.db.QueryRow(
		"INSERT INTO items (create_at, jsondata) VALUES ($1, $2) RETURNING id",
		modelObj.CreateAt, modelObj.JsonData,
	).Scan(&modelObj.Id)

	return err
}

func (i *ItemRepository) SelectItem() ([]byte, error) {
	defer i.storage.Close()
	if err := i.storage.Open(); err != nil {
		log.Fatal(err)
	}

	var jData []byte
	if err := i.storage.db.QueryRow(
		"SELECT jsondata FROM items LIMIT 1").Scan(&jData); err != nil {
		return nil, err
	}

	return jData, nil
}

func (i *ItemRepository) DropItem() error {
	defer i.storage.Close()
	if err := i.storage.Open(); err != nil {
		log.Fatal(err)
	}

	query, err := i.storage.db.Prepare("DELETE FROM items")
	if err != nil {
		return err
	}
	if _, err := query.Exec(); err != nil {
		return err
	}

	return nil
}
