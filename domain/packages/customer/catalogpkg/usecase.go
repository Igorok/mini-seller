package catalogpkg

import "mini-seller/domain/common/entities/catalogentity"

type IUseCase interface {
	GetOrganizationList() ([]catalogentity.OrganizationInfo, error)
	GetOrganizationInfo(id string) (catalogentity.OrganizationInfo, error)

	GetCategoryList() ([]catalogentity.CategoryInfo, error)
	GetCategoryDetail(id string) (catalogentity.CategoryInfo, error)

	GetProductList(id_organization string, id_category string) ([]catalogentity.ProductInfo, error)
	GetProductDetail(id string) (catalogentity.ProductInfo, error)
}
