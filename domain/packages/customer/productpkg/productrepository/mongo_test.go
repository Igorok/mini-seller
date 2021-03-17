package productrepository

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
		log.Fatal("Product", err)
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

func TestGetProductList(t *testing.T) {
	t.Log("Test product repository list")

	productRepo := NewProductRepository(db)

	// test skip limit working
	prodList, count, err := productRepo.GetProductList(context.TODO(), 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, count, int64(8))
	assert.Equal(t, len(prodList), 2)
	assert.Equal(t, prodList[0].Name, "Steak New York")
	assert.Equal(t, prodList[0].Category.Name, "Steak")
	assert.Equal(t, prodList[0].Organization.Name, "restaurant")

	prodList, count, err = productRepo.GetProductList(context.TODO(), 6, 2)
	assert.Nil(t, err)
	assert.Equal(t, count, int64(8))
	assert.Equal(t, len(prodList), 2)
	assert.Equal(t, prodList[0].Name, "Four Cheese")
	assert.Equal(t, prodList[0].Category.Name, "Pizza")
	assert.Equal(t, prodList[0].Organization.Name, "pizza")
}

func TestGetProductDetail(t *testing.T) {
	t.Log("Test product repository detail")

	productRepo := NewProductRepository(db)
	product, err := productRepo.GetProductDetail(context.TODO(), "604497558ffcad558eb8e1f6")

	assert.Nil(t, err)
	assert.Equal(t, product.Name, "Cola")
	assert.Equal(t, product.Category.Name, "Drinks")
	assert.Equal(t, product.Organization.Name, "pizza")
}
