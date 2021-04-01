package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/model"

	"github.com/prometheus/common/log"
)

func (r *categoryResolver) Products(ctx context.Context, obj *model.Category) ([]*model.Product, error) {
	IDOrg := ""
	if obj.IDOrg != "" {
		IDOrg = obj.IDOrg
	}

	prodList, err := r.CatalogUseCase.GetProductList(ctx, IDOrg, obj.ID)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if prodList == nil {
		return nil, nil
	}

	products := make([]*model.Product, len(prodList))
	for i, product := range prodList {
		products[i] = &model.Product{
			ID:             product.ID,
			IDCategory:     product.IDCategory,
			IDOrganization: product.IDOrganization,
			Name:           product.Name,
			Price:          product.Price,
			Count:          product.Count,
			Status:         product.Status,
		}
	}

	return products, nil
}

func (r *organizationResolver) Categories(ctx context.Context, obj *model.Organization) ([]*model.Category, error) {
	catList, err := r.CatalogUseCase.GetCategoryList(ctx, obj.IDsCategory)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if catList == nil {
		return nil, nil
	}

	categories := make([]*model.Category, len(catList))
	for i, category := range catList {
		categories[i] = &model.Category{
			ID:     category.ID,
			Name:   category.Name,
			Status: category.Status,
			IDOrg:  obj.ID,
		}
	}

	return categories, nil
}

func (r *organizationResolver) Products(ctx context.Context, obj *model.Organization) ([]*model.Product, error) {
	prodList, err := r.CatalogUseCase.GetProductList(ctx, obj.ID, "")
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if prodList == nil {
		return nil, nil
	}

	products := make([]*model.Product, len(prodList))
	for i, product := range prodList {
		products[i] = &model.Product{
			ID:             product.ID,
			IDCategory:     product.IDCategory,
			IDOrganization: product.IDOrganization,
			Name:           product.Name,
			Price:          product.Price,
			Count:          product.Count,
			Status:         product.Status,
		}
	}

	return products, nil
}

func (r *productResolver) Category(ctx context.Context, obj *model.Product) (*model.Category, error) {
	cat, err := r.CatalogUseCase.GetCategoryDetail(ctx, obj.IDCategory)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, nil
	}

	category := &model.Category{
		ID:     cat.ID,
		Name:   cat.Name,
		Status: cat.Status,
		IDOrg:  obj.IDOrganization,
	}

	return category, nil
}

func (r *productResolver) Organization(ctx context.Context, obj *model.Product) (*model.Organization, error) {
	org, err := r.CatalogUseCase.GetOrganizationDetail(ctx, obj.IDOrganization)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, nil
	}

	organization := &model.Organization{
		ID:          org.ID,
		Name:        org.Name,
		Email:       org.Email,
		Phone:       org.Phone,
		Status:      org.Status,
		IDsCategory: org.IDsCategory,
	}

	return organization, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type categoryResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
