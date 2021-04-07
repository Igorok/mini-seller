package catalogusecase

import (
	"context"
	"os"
	"testing"

	"mini-seller/domain/packages/catalogpkg/catalogrepository"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

// TestMain - integration tests for repository
func TestMain(m *testing.M) {
	os.Chdir("../../../../")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetOrganizationList(t *testing.T) {
	t.Log("Test catalogusecase organization list")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catalogUseCase := NewCatalogUseCase(catalogRepo)

	orgList, err := catalogUseCase.GetOrganizationList(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, len(orgList), 2)
	assert.Equal(t, orgList[0].Name, "pizza")
}

func TestGetOrganizationDetail(t *testing.T) {
	t.Log("Test catalogusecase organization detail")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catalogUseCase := NewCatalogUseCase(catalogRepo)

	organization, err := catalogUseCase.GetOrganizationDetail(context.TODO(), "6043d76e94df8de741c2c0d6")
	assert.Nil(t, err)
	assert.Equal(t, organization.Name, "restaurant")
}

func TestGetCategoryList(t *testing.T) {
	t.Log("Test catalogusecase categories list")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catalogUseCase := NewCatalogUseCase(catalogRepo)

	catList, err := catalogUseCase.GetCategoryList(context.TODO(), []string{"604488100f719d9c76a28fe3", "604488100f719d9c76a28fe4"})
	assert.Nil(t, err)
	assert.Equal(t, len(catList), 2)
	assert.Equal(t, catList[1].Name, "Alcohol")
}

func TestGetCategoryDetail(t *testing.T) {
	t.Log("Test catalogusecase category detail")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catalogUseCase := NewCatalogUseCase(catalogRepo)

	organization, err := catalogUseCase.GetCategoryDetail(context.TODO(), "604488100f719d9c76a28fe5")
	assert.Nil(t, err)
	assert.Equal(t, organization.Name, "Steak")
}

func TestGetProductList(t *testing.T) {
	t.Log("Test catalogusecase products list")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catalogUseCase := NewCatalogUseCase(catalogRepo)

	prodList, err := catalogUseCase.GetProductList(context.TODO(), nil, []string{"604488100f719d9c76a28fe3"})
	assert.Nil(t, err)
	assert.Equal(t, len(prodList), 2)
	assert.Equal(t, prodList[0].Name, "Cola")

	prodList, err = catalogUseCase.GetProductList(context.TODO(), []string{"6043d76e94df8de741c2c0d5"}, nil)
	assert.Nil(t, err)
	assert.Equal(t, len(prodList), 4)
	assert.Equal(t, prodList[0].Name, "Cola")

	prodList, err = catalogUseCase.GetProductList(context.TODO(), []string{"6043d76e94df8de741c2c0d5"}, []string{"604488100f719d9c76a28fe7"})
	assert.Nil(t, err)
	assert.Equal(t, len(prodList), 3)
	assert.Equal(t, prodList[0].Name, "Chicken Barbecue")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test catalogusecase product detail")

	catalogRepo := catalogrepository.NewCatalogRepositoryMock()
	catalogUseCase := NewCatalogUseCase(catalogRepo)

	product, err := catalogUseCase.GetProductDetail(context.TODO(), "604497558ffcad558eb8e1f4")
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Salad Cesar")
}
