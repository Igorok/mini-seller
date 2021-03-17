package productpkg

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
)

// IRepository - access to data
type IRepository interface {
	GetProductList(ctx context.Context, skip, limit int) ([]*productentity.ProductForList, int64, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.ProductForList, error)
}
