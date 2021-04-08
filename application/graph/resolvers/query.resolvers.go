package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/model"

	"github.com/prometheus/common/log"
)

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

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	prod, err := r.CatalogUseCase.GetProductDetail(ctx, id)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if prod == nil {
		return nil, nil
	}

	product := &model.Product{
		ID:             prod.ID,
		IDCategory:     prod.IDCategory,
		IDOrganization: prod.IDOrganization,
		Name:           prod.Name,
		Price:          prod.Price,
		Count:          prod.Count,
		Status:         prod.Status,
	}

	return product, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
