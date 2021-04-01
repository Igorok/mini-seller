package productrepository

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/productpkg"
	"mini-seller/infrastructure/mongohelper/testdata"
)

// RepositoryMock - mock instead database
type RepositoryMock struct {
	productList []*productentity.ProductForList
}

// NewProductRepositoryMock - constructor for mock of repository with fixed data
func NewProductRepositoryMock() *RepositoryMock {
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

	productList := make([]*productentity.ProductForList, len(prodData.Products))
	for i, prod := range prodData.Products {
		productList[i] = &productentity.ProductForList{
			ID:     prod.ID,
			Name:   prod.Name,
			Price:  prod.Price,
			Count:  prod.Count,
			Status: prod.Status,
			Category: productentity.CategoryForProduct{
				ID:   prod.IDCategory,
				Name: catById[prod.IDCategory].Name,
			},
			Organization: productentity.OrganizationForProduct{
				ID:    prod.IDOrganization,
				Name:  orgById[prod.IDOrganization].Name,
				Phone: orgById[prod.IDOrganization].Phone,
				Email: orgById[prod.IDOrganization].Email,
			},
		}
	}

	return &RepositoryMock{productList: productList}
}

// GetProductList - mock for products list
func (rm RepositoryMock) GetProductList(ctx context.Context, skip, limit int) ([]*productentity.ProductForList, int64, error) {
	prodList := rm.productList[skip : skip+limit]
	return prodList, int64(len(rm.productList)), nil
}

// GetProductList - mock for product detail
func (rm RepositoryMock) GetProductDetail(ctx context.Context, id string) (*productentity.ProductForList, error) {
	for _, product := range rm.productList {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, productpkg.ErrProductNotFound
}
