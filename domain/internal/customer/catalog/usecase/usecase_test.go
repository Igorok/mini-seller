package usecase

import (
	"context"
	"mini-seller/domain/internal/customer/catalog"
	"mini-seller/domain/internal/customer/catalog/repository"
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

	catalogRepo := repository.NewCatalogRepositoryMock()
	catUseCase := NewUseCase(catalogRepo)

	// test validation
	catalogProd, count, err := catUseCase.GetCatalog(context.TODO(), 0, 1000)
	assert.Nil(t, catalogProd)
	assert.Equal(t, err, catalog.ErrCatalogLimit)

	catalogProd, count, err = catUseCase.GetCatalog(context.TODO(), 0, 0)
	assert.Nil(t, catalogProd)
	assert.Equal(t, err, catalog.ErrCatalogLimit)

	// test use case
	catalogProd, count, err = catUseCase.GetCatalog(context.TODO(), 2, 2)
	assert.Nil(t, err)
	assert.Equal(t, count, int64(8))
	assert.Equal(t, len(catalogProd), 2)
	assert.Equal(t, catalogProd[0].Name, "Salad Cesar")
	assert.Equal(t, catalogProd[0].Category.Name, "Salad")
	assert.Equal(t, catalogProd[0].Organization.Name, "restaurant")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test product usecase detail")

	catalogRepo := repository.NewCatalogRepositoryMock()
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
