package catalogrepository

import (
	"context"
	"log"
	"mini-seller/infrastructure/mongohelper"
	"mini-seller/infrastructure/mongohelper/testdata"
	"mini-seller/infrastructure/viperhelper"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

// TestMain - integration tests for repository
func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	os.Chdir("../../../../")
	vip := viperhelper.Viper{ConfigType: "", ConfigName: "test.json", ConfigPath: "infrastructure/viperhelper"}
	err := vip.Read()
	if err != nil {
		log.Fatal("Catalog", err)
	}

	db, err = mongohelper.Connect("test_catalog")
	if err != nil {
		log.Fatal(err)
	}

	testdata.InsertOrganizations(db)
	testdata.InsertProducts(db)
}

func tearDown() {
	db.Drop(context.TODO())
}

func TestGetOrganizationList(t *testing.T) {
	t.Log("Test catalogrepository organization list")

	catalogRepo := NewCatalogRepository(db)

	orgList, err := catalogRepo.GetOrganizationList(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, len(orgList), 2)
	assert.Equal(t, orgList[0].Name, "pizza")
}

func TestGetOrganizationDetail(t *testing.T) {
	t.Log("Test catalogrepository organization detail")

	catalogRepo := NewCatalogRepository(db)

	organization, err := catalogRepo.GetOrganizationDetail(context.TODO(), "6043d76e94df8de741c2c0d6")
	assert.Nil(t, err)
	assert.Equal(t, organization.Name, "restaurant")
}

func TestGetCategoryList(t *testing.T) {
	t.Log("Test catalogrepository categories list")

	catalogRepo := NewCatalogRepository(db)
	catList, err := catalogRepo.GetCategoryList(context.TODO(), []string{"604488100f719d9c76a28fe3", "604488100f719d9c76a28fe4"})
	assert.Nil(t, err)
	assert.Equal(t, len(catList), 2)
	assert.Equal(t, catList[1].Name, "Alcohol")
}

func TestGetCategoryDetail(t *testing.T) {
	t.Log("Test catalogrepository category detail")

	catalogRepo := NewCatalogRepository(db)

	organization, err := catalogRepo.GetCategoryDetail(context.TODO(), "604488100f719d9c76a28fe5")
	assert.Nil(t, err)
	assert.Equal(t, organization.Name, "Steak")
}

func TestGetProductList(t *testing.T) {
	t.Log("Test catalogrepository products list")

	catalogRepo := NewCatalogRepository(db)

	prodList, err := catalogRepo.GetProductList(context.TODO(), nil, []string{"604488100f719d9c76a28fe3"})
	assert.Nil(t, err)
	assert.Equal(t, len(prodList), 2)
	assert.Equal(t, prodList[0].Name, "Cola")

	prodList, err = catalogRepo.GetProductList(context.TODO(), []string{"6043d76e94df8de741c2c0d5"}, nil)
	assert.Nil(t, err)
	assert.Equal(t, len(prodList), 4)
	assert.Equal(t, prodList[0].Name, "Cola")

	prodList, err = catalogRepo.GetProductList(context.TODO(), []string{"6043d76e94df8de741c2c0d5"}, []string{"604488100f719d9c76a28fe7"})
	assert.Nil(t, err)
	assert.Equal(t, len(prodList), 3)
	assert.Equal(t, prodList[0].Name, "Chicken Barbecue")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test catalogrepository product detail")

	catalogRepo := NewCatalogRepository(db)

	product, err := catalogRepo.GetProductDetail(context.TODO(), "604497558ffcad558eb8e1f4")
	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Salad Cesar")
}
