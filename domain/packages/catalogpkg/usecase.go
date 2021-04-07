package catalogpkg

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
)

type IUseCase interface {
	GetOrganizationList(ctx context.Context) ([]*organizationentity.Organization, error)
	GetOrganizationDetail(ctx context.Context, id string) (*organizationentity.Organization, error)

	GetCategoryList(ctx context.Context, ids []string) ([]*productcategoryentity.ProductCategory, error)
	GetCategoryDetail(ctx context.Context, id string) (*productcategoryentity.ProductCategory, error)

	GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.Product, error)
}
