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
	os.Chdir("../../../../../")
	vip := viperhelper.Viper{ConfigType: "", ConfigName: "test.json", ConfigPath: "infrastructure/viperhelper"}
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

func TestGetOrganizationList(t *testing.T) {
	t.Log("Test catalogrepository organization list")

	catalogRepo := NewCatalogRepository(db)

	orgList, err := catalogRepo.GetOrganizationList()
	assert.Nil(t, err)
	assert.Equal(t, len(orgList), 2)
	assert.Equal(t, orgList[0].Name, "pizza")
}

func TestGetOrganizationDetail(t *testing.T) {
	t.Log("Test catalogrepository organization list")

	catalogRepo := NewCatalogRepository(db)

	organization, err := catalogRepo.GetOrganizationDetail("6043d76e94df8de741c2c0d6")
	assert.Nil(t, err)
	assert.Equal(t, organization.Name, "restaurant")
}
