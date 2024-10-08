package repositories

import (
	"go-gpt-task/models"
	"go-gpt-task/usecases"
)

var (
	_ usecases.DbRepository = &Database{}
)

type Database []models.Laptop

func (db *Database) Insert(laptop models.Laptop) {
	*db = append(*db, laptop)
}

func (db *Database) FindByID(id string) (models.Laptop, bool) {
	for _, lp := range *db {
		if lp.ID == id {
			return lp, true
		}
	}

	return models.Laptop{}, false
}
