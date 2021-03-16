package catalogusecase

import (
	"context"
	"mini-seller/domain/packages/customer/catalog"
	"mini-seller/domain/packages/customer/catalog/catalogrepository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Chdir("../../../../../")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetCatalog(t *testing.T) {
	t.Log("Test product usecase list")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catUseCase := NewUseCase(catalogRepo)

	// test validation
	catalogProd, err := catUseCase.GetCatalog(context.TODO(), 0, 1000)
	assert.Nil(t, catalogProd.Products)
	assert.Equal(t, err, catalog.ErrCatalogLimit)

	catalogProd, err = catUseCase.GetCatalog(context.TODO(), 0, 0)
	assert.Nil(t, catalogProd.Products)
	assert.Equal(t, err, catalog.ErrCatalogLimit)

	// test use case
	catalogProd, err = catUseCase.GetCatalog(context.TODO(), 2, 2)
	assert.Nil(t, err)
	assert.Equal(t, catalogProd.Count, int64(8))
	assert.Equal(t, len(catalogProd.Products), 2)
	assert.Equal(t, catalogProd.Products[0].Name, "Salad Cesar")
	assert.Equal(t, catalogProd.Products[0].Category.Name, "Salad")
	assert.Equal(t, catalogProd.Products[0].Organization.Name, "restaurant")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test product usecase detail")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catUseCase := NewUseCase(catalogRepo)

	product, err := catUseCase.GetProductDetail(context.TODO(), "")
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, err, catalog.ErrProductNotFound)

	product, err = catUseCase.GetProductDetail(context.TODO(), "604497558ffcad558eb8e1f6")
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Cola")
	assert.Equal(t, product.Category.Name, "Drinks")
	assert.Equal(t, product.Organization.Name, "pizza")
}
