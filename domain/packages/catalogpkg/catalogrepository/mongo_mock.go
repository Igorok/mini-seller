package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/infrastructure/mongohelper/testdata"
)

// RepositoryMock - mock instead database
type RepositoryMock struct {
	organizations []*organizationentity.Organization
	products      []*productentity.Product
	categories    []*productcategoryentity.ProductCategory
}

// NewCatalogRepositoryMock - constructor for mock of repository with fixed data
func NewCatalogRepositoryMock() *RepositoryMock {
	orgData, _ := testdata.GetOrganizations()
	prodData, _ := testdata.GetProducts()

	organizations := make([]*organizationentity.Organization, len(orgData.Organizations))
	for i := range orgData.Organizations {
		organizations[i] = &orgData.Organizations[i]
	}

	products := make([]*productentity.Product, len(prodData.Products))
	for i := range prodData.Products {
		products[i] = &prodData.Products[i]
	}
	categories := make([]*productcategoryentity.ProductCategory, len(prodData.Categories))
	for i := range prodData.Categories {
		categories[i] = &prodData.Categories[i]
	}

	return &RepositoryMock{
		organizations: organizations,
		products:      products,
		categories:    categories,
	}
}

func includes(arr []string, val string) bool {
	for _, str := range arr {
		if str == val {
			return true
		}
	}
	return false
}

func (r *RepositoryMock) GetOrganizationList(ctx context.Context) ([]*organizationentity.Organization, error) {
	return r.organizations, nil
}

func (r *RepositoryMock) GetOrganizationDetail(ctx context.Context, id string) (*organizationentity.Organization, error) {
	for _, org := range r.organizations {
		if id == org.ID {
			return org, nil
		}
	}
	return nil, nil
}

func (r *RepositoryMock) GetCategoryList(ctx context.Context, ids []string) ([]*productcategoryentity.ProductCategory, error) {
	categories := make([]*productcategoryentity.ProductCategory, 0)

	for _, cat := range r.categories {
		if includes(ids, cat.ID) {
			categories = append(categories, cat)
		}
	}

	return categories, nil
}

func (r *RepositoryMock) GetCategoryDetail(ctx context.Context, id string) (*productcategoryentity.ProductCategory, error) {
	for _, cat := range r.categories {
		if id == cat.ID {
			return cat, nil
		}
	}
	return nil, nil
}

func prodQuery(ids_organization []string, ids_category []string, prod *productentity.Product) bool {
	var byOrg bool = true
	var byCat bool = true

	if ids_organization != nil {
		byOrg = includes(ids_organization, prod.IDOrganization)
	}

	if ids_category != nil {
		byCat = includes(ids_category, prod.IDCategory)
	}

	return byOrg && byCat
}

func (r *RepositoryMock) GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error) {
	products := make([]*productentity.Product, 0)

	for _, prod := range r.products {
		if prodQuery(ids_organization, ids_category, prod) {
			products = append(products, prod)
		}
	}

	return products, nil
}

func (r *RepositoryMock) GetProductDetail(ctx context.Context, id string) (*productentity.Product, error) {
	for _, prod := range r.products {
		if id == prod.ID {
			return prod, nil
		}
	}
	return nil, nil
}
