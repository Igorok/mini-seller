package catalogusecase

import (
	"context"
	"mini-seller/domain/common/entities/catalogentity"
	"mini-seller/domain/packages/catalogpkg/catalogrepository"
)

type UseCase struct {
	catalogRepo *catalogrepository.Repository
}

func NewCatalogUseCase(catalogRepo *catalogrepository.Repository) *UseCase {
	return &UseCase{catalogRepo: catalogRepo}
}

func (cUseCase UseCase) GetOrganizationList(ctx context.Context) ([]*catalogentity.OrganizationInfo, error) {
	return cUseCase.catalogRepo.GetOrganizationList(ctx)
}
func (cUseCase UseCase) GetOrganizationDetail(ctx context.Context, id string) (*catalogentity.OrganizationInfo, error) {
	return cUseCase.catalogRepo.GetOrganizationDetail(ctx, id)
}

func (cUseCase UseCase) GetCategoryList(ctx context.Context, ids []string) ([]*catalogentity.CategoryInfo, error) {
	return cUseCase.catalogRepo.GetCategoryList(ctx, ids)
}
func (cUseCase UseCase) GetCategoryDetail(ctx context.Context, id string) (*catalogentity.CategoryInfo, error) {
	return cUseCase.catalogRepo.GetCategoryDetail(ctx, id)
}

func (cUseCase UseCase) GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*catalogentity.ProductInfo, error) {
	return cUseCase.catalogRepo.GetProductList(ctx, ids_organization, ids_category)
}
func (cUseCase UseCase) GetProductDetail(ctx context.Context, id string) (*catalogentity.ProductInfo, error) {
	return cUseCase.catalogRepo.GetProductDetail(ctx, id)
}
