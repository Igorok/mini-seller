package catalog

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
)

// IRepository - access to data
type IRepository interface {
	GetCatalog(ctx context.Context, skip, limit int) ([]*productentity.ProductForCatalog, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.ProductForCatalog, error)
}
