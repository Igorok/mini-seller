package catalogpkg

import (
	"context"
	"mini-seller/domain/common/entities/catalogentity"
)

type IUseCase interface {
	GetOrganizationList(ctx context.Context) ([]*catalogentity.OrganizationInfo, error)
	GetOrganizationDetail(ctx context.Context, id string) (*catalogentity.OrganizationInfo, error)

	GetCategoryList(ctx context.Context, ids []string) ([]*catalogentity.CategoryInfo, error)
	GetCategoryDetail(ctx context.Context, id string) (*catalogentity.CategoryInfo, error)

	GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*catalogentity.ProductInfo, error)
	GetProductDetail(ctx context.Context, id string) (*catalogentity.ProductInfo, error)
}
