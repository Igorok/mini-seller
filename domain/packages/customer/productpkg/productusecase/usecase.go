package productusecase

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/customer/productpkg"
)

// UseCase - use case for business logic
type UseCase struct {
	repository productpkg.IRepository
}

// NewUseCase - constructor for use case
func NewUseCase(repository productpkg.IRepository) UseCase {
	return UseCase{repository: repository}
}

// GetProductList - list of products
func (uc UseCase) GetProductList(ctx context.Context, skip, limit int) (*productentity.ProductList, error) {
	if limit == 0 || limit > 100 {
		return &productentity.ProductList{}, productpkg.ErrListLimit
	}

	products, count, err := uc.repository.GetProductList(ctx, skip, limit)
	if err != nil {
		return &productentity.ProductList{}, err
	}

	return &productentity.ProductList{Products: products, Count: count}, nil
}

// GetProductDetail - details for products
func (uc UseCase) GetProductDetail(ctx context.Context, id string) (*productentity.ProductForList, error) {
	if id == "" {
		return nil, productpkg.ErrProductNotFound
	}
	return uc.repository.GetProductDetail(ctx, id)
}
