package catalog

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
	os.Chdir("../../../../../")
	vip := viperhelper.Viper{ConfigType: "", ConfigName: "test", ConfigPath: "infrastructure/viperhelper"}
	err := vip.Read()
	if err != nil {
		log.Fatal("Catalog", err)
	}

	db, err = mongohelper.Connect()
	if err != nil {
		log.Fatal(err)
	}

	testdata.InsertOrganizations(db)
	testdata.InsertProducts(db)
}

func tearDown() {
	db.Drop(context.TODO())
}

func TestGetCatalog(t *testing.T) {
	t.Log("Test product list")

	catalogRepo := NewCatalogRepository(db)
	catalogProd, err := catalogRepo.GetCatalog(context.TODO(), 0, 10)

	assert.Nil(t, err)
	assert.Equal(t, len(catalogProd), 8)
	assert.Equal(t, catalogProd[0].Name, "Steak New York")
	assert.Equal(t, catalogProd[0].Category.Name, "Steak")
	assert.Equal(t, catalogProd[0].Organization.Name, "restaurant")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test product detail")

	catalogRepo := NewCatalogRepository(db)
	product, err := catalogRepo.GetProductDetail(context.TODO(), "604497558ffcad558eb8e1f6")

	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Cola")
	assert.Equal(t, product.Category.Name, "Drinks")
	assert.Equal(t, product.Organization.Name, "pizza")
}
