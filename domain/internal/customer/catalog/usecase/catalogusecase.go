package catalog

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/internal/customer/catalog"
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
func (uc UseCase) GetCatalog(ctx context.Context, skip, limit int) ([]*productentity.ProductForCatalog, error) {
	return uc.repository.GetCatalog(ctx, skip, limit)
}

// GetProductDetail - details for products
func (uc UseCase) GetProductDetail(ctx context.Context, id string) (*productentity.ProductForCatalog, error) {
	if id == "" {
		return nil, catalog.ErrProductNotFound
	}
	return uc.repository.GetProductDetail(ctx, id)
}
