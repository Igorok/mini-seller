package repository

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/internal/customer/catalog"
	"mini-seller/infrastructure/mongohelper/testdata"
)

// RepositoryMock - mock instead database
type RepositoryMock struct {
	productList []*productentity.ProductForCatalog
}

// NewCatalogRepositoryMock - constructor for mock of repository with fixed data
func NewCatalogRepositoryMock() *RepositoryMock {
	orgData, _ := testdata.GetOrganizations()
	prodData, _ := testdata.GetProducts()

	orgById := make(map[string]organizationentity.Organization)
	catById := make(map[string]productcategoryentity.ProductCategory)

	for _, org := range orgData.Organizations {
		orgById[org.ID] = org
	}
	for _, cat := range prodData.Categories {
		catById[cat.ID] = cat
	}

	productList := make([]*productentity.ProductForCatalog, len(prodData.Products))
	for i, prod := range prodData.Products {
		productList[i] = &productentity.ProductForCatalog{
			ID:     prod.ID,
			Name:   prod.Name,
			Price:  prod.Price,
			Count:  prod.Count,
			Status: prod.Status,
			Category: productentity.CategoryForCatalog{
				ID:   prod.IDCategory,
				Name: catById[prod.IDCategory].Name,
			},
			Organization: productentity.OrganizationForCatalog{
				ID:    prod.IDOrganization,
				Name:  orgById[prod.IDOrganization].Name,
				Phone: orgById[prod.IDOrganization].Phone,
				Email: orgById[prod.IDOrganization].Email,
			},
		}
	}

	return &RepositoryMock{productList: productList}
}

// GetCatalog - mock for products list
func (rm RepositoryMock) GetCatalog(ctx context.Context, skip, limit int) ([]*productentity.ProductForCatalog, int64, error) {
	prodList := rm.productList[skip : skip+limit]
	return prodList, int64(len(rm.productList)), nil
}

// GetCatalog - mock for product detail
func (rm RepositoryMock) GetProductDetail(ctx context.Context, id string) (*productentity.ProductForCatalog, error) {
	for _, product := range rm.productList {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, catalog.ErrProductNotFound
}
