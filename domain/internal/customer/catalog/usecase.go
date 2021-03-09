package catalog

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
)

// IUseCase - use case for business logic
type IUseCase interface {
	GetCatalog(ctx context.Context, skip, limit int) ([]*productentity.ProductForCatalog, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.ProductForCatalog, error)
}
