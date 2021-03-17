package productusecase

import (
	"context"
	"mini-seller/domain/packages/customer/productpkg"
	"mini-seller/domain/packages/customer/productpkg/productrepository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Chdir("../../../../../")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetProductList(t *testing.T) {
	t.Log("Test product usecase list")

	productRepo := productrepository.NewProductRepositoryMock()
	productUseCase := NewUseCase(productRepo)

	// test validation
	productList, err := productUseCase.GetProductList(context.TODO(), 0, 1000)
	assert.Nil(t, productList.Products)
	assert.Equal(t, err, productpkg.ErrListLimit)

	productList, err = productUseCase.GetProductList(context.TODO(), 0, 0)
	assert.Nil(t, productList.Products)
	assert.Equal(t, err, productpkg.ErrListLimit)

	// test use case
	productList, err = productUseCase.GetProductList(context.TODO(), 2, 2)
	assert.Nil(t, err)
	assert.Equal(t, productList.Count, int64(8))
	assert.Equal(t, len(productList.Products), 2)
	assert.Equal(t, productList.Products[0].Name, "Salad Cesar")
	assert.Equal(t, productList.Products[0].Category.Name, "Salad")
	assert.Equal(t, productList.Products[0].Organization.Name, "restaurant")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test product usecase detail")

	productRepo := productrepository.NewProductRepositoryMock()
	productUseCase := NewUseCase(productRepo)

	prod, err := productUseCase.GetProductDetail(context.TODO(), "")
	assert.Nil(t, prod)
	assert.NotNil(t, err)
	assert.Equal(t, err, productpkg.ErrProductNotFound)

	prod, err = productUseCase.GetProductDetail(context.TODO(), "604497558ffcad558eb8e1f6")
	assert.Nil(t, err)
	assert.Equal(t, prod.Name, "Cola")
	assert.Equal(t, prod.Category.Name, "Drinks")
	assert.Equal(t, prod.Organization.Name, "pizza")
}
