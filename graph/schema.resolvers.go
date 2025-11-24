package graph

import (
	"context"

	"github.com/gabezy/go-graphql-api/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindByCategoryID(obj.ID)
	if err != nil {
		return nil, err
	}

	coursesModel := model.DBListToModelList(courses)
	return coursesModel, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input *model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	cat, err := r.CategoryDB.FindByID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	categoryModel := model.Category{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: &cat.Description,
	}

	newCourse, err := r.CourseDB.Create(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:          newCourse.ID,
		Name:        newCourse.Name,
		Description: &newCourse.Description,
		Category:    &categoryModel,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categoriesDB, err := r.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []*model.Category

	for _, cat := range categoriesDB {
		categories = append(categories, &model.Category{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: &cat.Description,
		})
	}

	return categories, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}

	var coursesModel []*model.Course

	for _, course := range courses {
		category, err := r.CategoryDB.FindByID(course.CategoryID)
		if err != nil {
			return nil, err
		}

		categoryModel := model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		}

		coursesModel = append(coursesModel, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: &course.Description,
			Category:    &categoryModel,
		})
	}

	return coursesModel, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	categoryResolver struct{ *Resolver }
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
