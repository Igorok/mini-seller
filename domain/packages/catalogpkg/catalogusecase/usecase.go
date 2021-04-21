package catalogusecase

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/catalogpkg"
)

type UseCase struct {
	catalogRepo catalogpkg.IRepository
}

func NewCatalogUseCase(catalogRepo catalogpkg.IRepository) *UseCase {
	return &UseCase{catalogRepo: catalogRepo}
}

func (cUseCase UseCase) GetOrganizationList(ctx context.Context) ([]*organizationentity.Organization, error) {
	return cUseCase.catalogRepo.GetOrganizationList(ctx)
}
func (cUseCase UseCase) GetOrganizationDetail(ctx context.Context, id string) (*organizationentity.Organization, error) {
	if id == "" {
		return nil, catalogpkg.ErrOrganizationNotFound
	}
	return cUseCase.catalogRepo.GetOrganizationDetail(ctx, id)
}

func (cUseCase UseCase) GetCategoryList(ctx context.Context, ids []string) ([]*productcategoryentity.ProductCategory, error) {
	if ids == nil {
		return nil, catalogpkg.ErrCategoryNotFound
	}
	return cUseCase.catalogRepo.GetCategoryList(ctx, ids)
}
func (cUseCase UseCase) GetCategoryDetail(ctx context.Context, id string) (*productcategoryentity.ProductCategory, error) {
	if id == "" {
		return nil, catalogpkg.ErrCategoryNotFound
	}
	return cUseCase.catalogRepo.GetCategoryDetail(ctx, id)
}

func (cUseCase UseCase) GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error) {
	return cUseCase.catalogRepo.GetProductList(ctx, ids_organization, ids_category)
}
func (cUseCase UseCase) GetProductDetail(ctx context.Context, id string) (*productentity.Product, error) {
	if id == "" {
		return nil, catalogpkg.ErrProductNotFound
	}
	return cUseCase.catalogRepo.GetProductDetail(ctx, id)
}
