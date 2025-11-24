package model

import "github.com/gabezy/go-graphql-api/internal/database"

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Category    *Category `json:"category"`
}

func DBToModel(db database.Course) Course {
	return Course{
		ID:          db.ID,
		Name:        db.Name,
		Description: &db.Description,
	}
}

func DBListToModelList(dbModels []database.Course) []*Course {
	var courses []*Course
	for _, course := range dbModels {
		courseModel := DBToModel(course)
		courses = append(courses, &courseModel)
	}
	return courses
}
