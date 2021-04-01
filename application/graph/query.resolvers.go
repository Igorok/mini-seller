package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/model"

	"github.com/prometheus/common/log"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		// User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Organizations(ctx context.Context) ([]*model.Organization, error) {
	orgs, err := r.CatalogUseCase.GetOrganizationList(ctx)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if len(orgs) == 0 {
		return nil, nil
	}

	organizations := make([]*model.Organization, len(orgs))
	for i, org := range orgs {
		organizations[i] = &model.Organization{
			ID:          org.ID,
			Name:        org.Name,
			Email:       org.Email,
			Phone:       org.Phone,
			Status:      org.Status,
			IDsCategory: org.IDsCategory,
		}
	}

	return organizations, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
