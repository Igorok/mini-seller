package productpkg

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
)

// IUseCase - use case for business logic
type IUseCase interface {
	GetProductList(ctx context.Context, skip, limit int) (*productentity.ProductList, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.ProductForList, error)
}
