package model

import "github.com/gabezy/go-graphql-api/internal/database"

type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

func CategoryDBToModel(categoryDB database.Category) Category {
	return Category{
		ID:          categoryDB.ID,
		Name:        categoryDB.Name,
		Description: &categoryDB.Description,
	}
}

func CategoryDBListToModelList(categoriesDB []database.Category) []*Category {
	var categories []*Category
	for _, category := range categoriesDB {
		categoryModel := CategoryDBToModel(category)
		categories = append(categories, &categoryModel)
	}
	return categories
}
