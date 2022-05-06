package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/lkeix/graphql-sandbox/graph/generated"
	"github.com/lkeix/graphql-sandbox/graph/model"
)

const (
	layout = "2006-01-02 15:04:05"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	timestamp := time.Now().Format(layout)

	task := model.Task{
		Title:     input.Title,
		Note:      input.Note,
		Completed: 0,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	r.DB.Create(&task)

	return &task, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
