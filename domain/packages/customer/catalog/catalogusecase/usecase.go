package catalogusecase

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/customer/catalog"
)

// UseCase - use case for business logic
type UseCase struct {
	repository catalog.IRepository
}

// NewUseCase - constructor for use case
func NewUseCase(repository catalog.IRepository) UseCase {
	return UseCase{repository: repository}
}

// GetCatalog - list of products
func (uc UseCase) GetCatalog(ctx context.Context, skip, limit int) (*productentity.CatalogData, error) {
	if limit == 0 || limit > 100 {
		return &productentity.CatalogData{}, catalog.ErrCatalogLimit
	}

	products, count, err := uc.repository.GetCatalog(ctx, skip, limit)
	if err != nil {
		return &productentity.CatalogData{}, err
	}

	return &productentity.CatalogData{Products: products, Count: count}, nil
}

// GetProductDetail - details for products
func (uc UseCase) GetProductDetail(ctx context.Context, id string) (*productentity.ProductForCatalog, error) {
	if id == "" {
		return nil, catalog.ErrProductNotFound
	}
	return uc.repository.GetProductDetail(ctx, id)
}
