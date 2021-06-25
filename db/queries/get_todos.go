package queries

import (
	"clk/db"
	"clk/db/models"
)

func GetToDos(showInActive bool, page, limit int) ([]models.ToDo, error) {

	var todos []models.ToDo
	query := db.Client.
		Order("created_at desc").
		Offset(page * limit)

	if limit > 0 {
		query = query.Limit(limit)
	}

	if !showInActive {
		query = query.Where("active = 1")
	}

	result := query.Find(&todos)
	if result.Error != nil {
		return todos, result.Error
	}

	return todos, nil
}
