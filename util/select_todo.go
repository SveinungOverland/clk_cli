package util

import (
	"clk/db/models"
	"clk/db/queries"

	"github.com/manifoldco/promptui"
)

func SelectToDo(showInactive bool) (todo models.ToDo, err error) {
	var page int
	const limit = 10
	var items []interface{}
	var itemsLength int

SELECT:
	todos, err := queries.GetToDos(showInactive, page, limit)
	if err != nil {
		return
	}

	// BEGIN - PAGE Handling
	itemsLength = len(todos)
	if page > 0 {
		itemsLength++
	}
	if len(todos) == limit {
		itemsLength++
	}
	items = make([]interface{}, itemsLength)
	for i, todo := range todos {
		items[i] = todo
	}
	if len(todos) == limit {
		items[itemsLength-2] = "Next page"
	}
	if page > 0 {
		items[itemsLength-1] = "Previous page"
	}
	// END - PAGE Handling

	prompt := promptui.Select{
		Label: "Select todo",
		Items: items,
	}

	index, _, err := prompt.Run()

	if err != nil {
		return
	}

	if len(todos) == limit && index == itemsLength-1 {
		page++
		goto SELECT
	} else if page > 0 && index == itemsLength {
		page--
		goto SELECT
	}

	todo = todos[index]
	return
}
