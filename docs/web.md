## Web api with golang for beginners
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software etc. I want to build web api with golang, because big part of my work is backend of web or mobile applications.


### GVM
First step for work with golang is install golang.
Gvm - golang version manager, it's util to install and use different versions of golang.
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

gvm install go1.15.6 -B
gvm use go1.15.6 [--default]
gvm list
```

### Go modules
If i want to make workable project with any language i should install libraries from other developers and organizations. Golang have stuff called Go Modules - a module is a collection of Go packages stored in a file tree with a go.mod file at its root.
```
go mod init mini-seller
```

### Viper
When i develop programing product i should have configuration for web server, databases, file storage, api for integrations. Viper is comfortable and powerful package for management of configuration.
```
go get github.com/spf13/viper
```

### Mongo
To save data for my project i will use Mongo it's nosql database with simple syntax, big possibilities and good performance.
```
go get go.mongodb.org/mongo-driver
```

### Testing
When i develop something bigger than landing page i should know that my functionality is working. Launch tests in golang is very simple.
```
go test -v ./...
```

### Gqlgen
And one of important things for web application is itself web application. I select gqlgen for this, it have generation of code and big possibilities to make graphql api.
```
go get github.com/99designs/gqlgen
```

### Desing
After installation of packages need to plane design of application. Building a structure of big application is not simple, and before do it, very helpful to read about solid rules, clean architecture, n-tier architecture. Very shortly central part of application is entities of business logic. Business logic of project builds around entities. Logic should not be depened from web frameworks or data storages like orm or web api. Use cases shouldn't receive data from storages but  should use classes of repository for this. Repositories will help to change one database to another or message broker or web api. For testing of business logic uses mocks of repositories. If you have big difficult validation for your logic you should put this in classes of specifications. And you should not relate your logic with web framework, instead of this web application should depended from use cases, this make changing of web framework more simple.

In my case i made folders:

1. application - code for web application
    1. server - web server for project
    2. graph - folder with gqlgen application
        1. schemas - folder contain schemas for graphql
        2. model - folder contain entities for graphql
        3. resolvers - folder contain controllers for graphql

2. domain - business logic of project
    1. common - common data for all packages
        1. entities - here i save all entities, these could be entities for business logic, models for database and dto for communication between classes
    2. packages - here logic of project
        1. catalogpkg - logic for catalog of products
            1. usecase - interface for use case, it describe functionality available in package
            2. repository - interface describe functionality of data storage
            3. catalogusecase - contain use case - business logic of catalog and test for use case
            4. catalogrepository - contain repository with requests for database, integration tests for validation of repository, and mock of repository, it needed to test logic of use case without real connection to database
3. infrastructure - code without business logic, database driver, helper for configuration
    1. mongohelper - contain functionality for configure connection to mongo database and folder with test data. These files need to initialize empty project with demo data, and i use these for testing of logic.
    2. viperhelper - i use this helper to read configuration for project and for test. This helps me to change configuration from default to local and use environment variables.

### Infrastructure
Viper helper read local configuration for project and tests, and can use environment variables instead of configuration files.
```
// Package viperhelper - helper for viper, uses to apply configuration from local files and environment
package viperhelper

import (
	"os"

	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

// IViper - helper for viper, describe methods for initiation of config
type IViper interface {
	updateSettings()
	Read() error
	getEnv()
}

// default values for json config
const (
	configType = "json"
	configName = "config.json"
	configPath = "./infrastructure/viperhelper"
)

// list of environment variables
var envVariables = []string{
	"MONGO_DB", "MONGO_HOST", "MONGO_PORT", "MONGO_USER",
	"MONGO_PASSWORD", "MONGO_AUTH", "MONGO_REPLICASET",
	"WEB_PORT",
}

// Viper - class for initialization of viper config with values from local configuration and environment
type Viper struct {
	ConfigType, ConfigName, ConfigPath string
}

// updateSettings - set default values for arguments
func (vip *Viper) updateSettings() {
	if vip.ConfigType == "" {
		vip.ConfigType = configType
	}
	if vip.ConfigName == "" {
		vip.ConfigName = configName
	}
	if vip.ConfigPath == "" {
		vip.ConfigPath = configPath
	}
}

// getEnv - get variables from environment
func (vip *Viper) getEnv() {
	for _, variable := range envVariables {
		value := os.Getenv(variable)
		if value != "" {
			viper.Set(variable, value)
		}
	}
}

// Read - read configuration
func (vip *Viper) Read() error {
	vip.updateSettings()

	viper.SetConfigType(vip.ConfigType)
	viper.AddConfigPath(vip.ConfigPath)

	viper.SetConfigName("local-" + vip.ConfigName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Info("Local config", err)
		viper.SetConfigName(vip.ConfigName)
		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	vip.getEnv()

	return nil
}
```

Mongo helper take configuration from viper and connect to database, with authentication and replicaset.
```
package mongohelper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect connection to mongodb
func Connect(dbname string) (*mongo.Database, error) {
	fmt.Println("init mongodb")
	if dbname == "" {
		dbname = viper.GetString("MONGO_DB")
	}

	// connection string
	uri := "mongodb://"
	if viper.GetBool("MONGO_AUTH") {
		uri += viper.GetString("MONGO_USER") + ":" + viper.GetString("MONGO_PASSWORD") + "@"
	}
	uri += viper.GetString("MONGO_HOST") + ":" + viper.GetString("MONGO_PORT")

	replicaSet := viper.GetString("MONGO_REPLICASET")
	if replicaSet != "" {
		uri += "/" + dbname + "?replicaSet=" + replicaSet
	}

	// create connection
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("err")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database(dbname), nil
}
```

Test data - mongodb helper, contain folder with default data for project. Data helper give possibility read default data, and insert this data to database. This helpful for launching empty project and testing.
```
package testdata

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mini-seller/domain/common/entities/employeeentity"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductsData - dto for products test data
type ProductsData struct {
	Categories []productcategoryentity.ProductCategory
	Products   []productentity.Product
}

// GetProducts - read json of products
func GetProducts() (*ProductsData, error) {
	content, err := ioutil.ReadFile("infrastructure/mongohelper/testdata/products.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var productData ProductsData
	err = json.Unmarshal(content, &productData)
	if err != nil {
		log.Fatal("Unmarshal: ", err)
		return nil, err
	}

	return &productData, nil
}

// InsertProducts - insert test data for categories and products
func InsertProducts(db *mongo.Database) {
	productData, err := GetProducts()
	if err != nil {
		log.Fatal("Unmarshal: ", err)
	}

	categories := make([]interface{}, len(productData.Categories))
	for i, cat := range productData.Categories {
		entity := productcategoryentity.ToMongo(&cat)
		categories[i] = entity
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Collection("product_categories").DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Collection("product_categories").InsertMany(ctx, categories)
	if err != nil {
		log.Fatal(err)
	}

	products := make([]interface{}, len(productData.Products))
	for i, product := range productData.Products {
		entity := productentity.ToProductMongo(&product)
		products[i] = entity
	}

	_, err = db.Collection("products").DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Collection("products").InsertMany(ctx, products)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Entities
Entities this is human understandable entities of business logic like product or customer.
```
package productentity

// Product - entity of product
type Product struct {
	ID             string
	IDCategory     string
	IDOrganization string
	Name           string
	Price          int
	Count          int
	Status         string
}

package productcategoryentity

// ProductCategory is the entity of category
type ProductCategory struct {
	ID     string
	Name   string
	Status string
}

package organizationentity

// Organization is the entity of organization
type Organization struct {
	ID          string   `validate:"omitempty"`
	Name        string   `validate:"required"`
	Email       string   `validate:"required"`
	Phone       string   `validate:"required"`
	Status      string   `validate:"required"`
	IDsCategory []string `validate:"required"`
}
```

### Models
Models contain realization of entity for saving in database and methods to convert entity to model. At example shows model for product.
```
package productentity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductMongo - model for mongo database
type ProductMongo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	IDCategory     primitive.ObjectID `bson:"_id_cat"`
	IDOrganization primitive.ObjectID `bson:"_id_org"`
	Name           string
	Price          int
	Count          int
	Status         string
}

// ToProduct - method to convert model to entity
func ToProduct(pm *ProductMongo) *Product {
	ID := ""
	if !pm.ID.IsZero() {
		ID = pm.ID.Hex()
	}

	return &Product{
		ID:             ID,
		IDCategory:     pm.IDCategory.Hex(),
		IDOrganization: pm.IDOrganization.Hex(),
		Name:           pm.Name,
		Price:          pm.Price,
		Count:          pm.Count,
		Status:         pm.Status,
	}
}

// ToProductMongo - method to convert entity to model
func ToProductMongo(p *Product) *ProductMongo {
	var ID primitive.ObjectID
	if p.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(p.ID)
	}
	idCategory, _ := primitive.ObjectIDFromHex(p.IDCategory)
	idOrganization, _ := primitive.ObjectIDFromHex(p.IDOrganization)

	return &ProductMongo{
		ID:             ID,
		IDCategory:     idCategory,
		IDOrganization: idOrganization,
		Name:           p.Name,
		Price:          p.Price,
		Count:          p.Count,
		Status:         p.Status,
	}
}
```

### Interfaces
Interfaces describe classes. Communication between packages should be abstract and builds by interfaces. It help change realisation of methods and apply test classes instead real.
```
package catalogpkg

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
)

type IUseCase interface {
	GetOrganizationList(ctx context.Context) ([]*organizationentity.Organization, error)
	GetOrganizationDetail(ctx context.Context, id string) (*organizationentity.Organization, error)

	GetCategoryList(ctx context.Context, ids []string) ([]*productcategoryentity.ProductCategory, error)
	GetCategoryDetail(ctx context.Context, id string) (*productcategoryentity.ProductCategory, error)

	GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.Product, error)
}

type IRepository interface {
	GetOrganizationList(ctx context.Context) ([]*organizationentity.Organization, error)
	GetOrganizationDetail(ctx context.Context, id string) (*organizationentity.Organization, error)

	GetCategoryList(ctx context.Context, ids []string) ([]*productcategoryentity.ProductCategory, error)
	GetCategoryDetail(ctx context.Context, id string) (*productcategoryentity.ProductCategory, error)

	GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error)
	GetProductDetail(ctx context.Context, id string) (*productentity.Product, error)
}
```

### Use case
Use cases contain logic of application. Example contain simple validation and calling of repositories for api of products.

```
package catalogusecase

import (
	"context"
	"mini-seller/domain/common/entities/organizationentity"
	"mini-seller/domain/common/entities/productcategoryentity"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/catalogpkg"
)

type UseCase struct {
	catalogRepo catalogpkg.IRepository
}

func NewCatalogUseCase(catalogRepo catalogpkg.IRepository) *UseCase {
	return &UseCase{catalogRepo: catalogRepo}
}

func (cUseCase UseCase) GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error) {
	return cUseCase.catalogRepo.GetProductList(ctx, ids_organization, ids_category)
}
func (cUseCase UseCase) GetProductDetail(ctx context.Context, id string) (*productentity.Product, error) {
	if id == "" {
		return nil, catalogpkg.ErrProductNotFound
	}
	return cUseCase.catalogRepo.GetProductDetail(ctx, id)
}
```

### Test for use case
Tests contain verification of use case behavior with valid data and invalid data. Example for products.

```
package catalogusecase

import (
	"context"
	"os"
	"testing"

	"mini-seller/domain/packages/catalogpkg"
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

	product, err = catalogUseCase.GetProductDetail(context.TODO(), "")
	assert.Equal(t, err, catalogpkg.ErrProductNotFound)
}
```

### Repository
Repositories contain methods for management of data. Example for search products in mongodb.

```
package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/domain/packages/catalogpkg"
	"time"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
}

// NewCatalogRepository - constructor for catalog repository
func NewCatalogRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (cRepo Repository) GetProductList(contx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// query
	query := bson.M{"status": catalogpkg.StatusActive}

	if ids_organization != nil && len(ids_organization) > 0 {
		ids_org := make([]primitive.ObjectID, len(ids_organization))
		for i, id := range ids_organization {
			id_org, err_id := primitive.ObjectIDFromHex(id)
			if err_id != nil {
				return nil, err_id
			}
			ids_org[i] = id_org
		}
		query["_id_org"] = bson.M{"$in": ids_org}
	}

	if ids_category != nil && len(ids_category) > 0 {
		ids_cat := make([]primitive.ObjectID, len(ids_category))
		for i, id := range ids_category {
			id_cat, err_id := primitive.ObjectIDFromHex(id)
			if err_id != nil {
				return nil, err_id
			}
			ids_cat[i] = id_cat
		}
		query["_id_cat"] = bson.M{"$in": ids_cat}
	}

	// sort
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	// request
	cursor, err := cRepo.db.Collection("products").Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}

	// format data from cursor
	products := make([]*productentity.Product, 0)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		prodMongo := productentity.ProductMongo{}
		err := cursor.Decode(&prodMongo)
		if err != nil {
			return nil, err
		}
		product := productentity.ToProduct(&prodMongo)
		products = append(products, product)
	}

	// answer
	return products, nil
}

func (cRepo Repository) GetProductDetail(contx context.Context, id string) (*productentity.Product, error) {
	// convert id to bson
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// get context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// context
	prodMongo := productentity.ProductMongo{}
	err = cRepo.db.Collection("products").FindOne(ctx, bson.M{"_id": ID, "status": catalogpkg.StatusActive}).Decode(&prodMongo)
	if err != nil {
		return nil, err
	}

	// convert to entity
	product := productentity.ToProduct(&prodMongo)

	// answer
	return product, nil
}
```

### Tests for Repository
Tests of repositories should check that database queries in repository is valid. To check this, i insert default data into test database, launch repository and compare answer with correct result.
```
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
```

### Mock for Repository
Mock of repository is necessary for testing of use cases. Business logic should not depend from data source, it works with entities.
```
package catalogrepository

import (
	"context"
	"mini-seller/domain/common/entities/productentity"
	"mini-seller/infrastructure/mongohelper/testdata"
)

// RepositoryMock - mock instead database
type RepositoryMock struct {
	products      []*productentity.Product
}

// NewCatalogRepositoryMock - constructor for mock of repository with fixed data
func NewCatalogRepositoryMock() *RepositoryMock {
	prodData, _ := testdata.GetProducts()

	products := make([]*productentity.Product, len(prodData.Products))
	for i := range prodData.Products {
		products[i] = &prodData.Products[i]
	}
	categories := make([]*productcategoryentity.ProductCategory, len(prodData.Categories))
	for i := range prodData.Categories {
		categories[i] = &prodData.Categories[i]
	}

	return &RepositoryMock{
		organizations: organizations,
		products:      products,
		categories:    categories,
	}
}

func includes(arr []string, val string) bool {
	for _, str := range arr {
		if str == val {
			return true
		}
	}
	return false
}

func prodQuery(ids_organization []string, ids_category []string, prod *productentity.Product) bool {
	var byOrg bool = true
	var byCat bool = true

	if ids_organization != nil {
		byOrg = includes(ids_organization, prod.IDOrganization)
	}

	if ids_category != nil {
		byCat = includes(ids_category, prod.IDCategory)
	}

	return byOrg && byCat
}

func (r *RepositoryMock) GetProductList(ctx context.Context, ids_organization []string, ids_category []string) ([]*productentity.Product, error) {
	products := make([]*productentity.Product, 0)

	for _, prod := range r.products {
		if prodQuery(ids_organization, ids_category, prod) {
			products = append(products, prod)
		}
	}

	return products, nil
}

func (r *RepositoryMock) GetProductDetail(ctx context.Context, id string) (*productentity.Product, error) {
	for _, prod := range r.products {
		if id == prod.ID {
			return prod, nil
		}
	}
	return nil, nil
}
```

### Gqlgen
Initialization of gqlgen with recommended folder structure by running this command
```
go run github.com/99designs/gqlgen init
```

My structure of folders:
- schemas
    - catalog.graphqls
    - query.graphqls
- model
    - catalog
- resolvers
    - catalog.resolvers.go
    - query.resolvers.go
    - catalogdataloader.go
    - resolver.go
- generated


Gqlgen can generate models and code for resolvers by schemas
```
go run github.com/99designs/gqlgen generate
```

To run go generate recursively over your entire project, use this command:
```
go generate ./...
```

### Schemas
```
# entities
type Organization {
  id: String!
  name: String!
  email: String!
  phone: String!
  status: String!
  categories: [Category]
  products: [Product]
}

type Category {
  id: String!
  name: String!
  status: String!
  products: [Product]
}

type Product {
  id: String!
  name: String!
  status: String!
  price: Int!
  count: Int!
  category: Category
  organization: Organization
}

# query
type Query {
  organizations: [Organization!]!
  product(id: String!): Product
}
```

### Models
```
package model

// Product - entity for product
type Product struct {
	ID             string `json:"id"`
	IDCategory     string `json:"category"`
	IDOrganization string `json:"organization"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Count          int    `json:"count"`
	Status         string `json:"status"`
}

// Category - entity for category of products
type Category struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	IDOrg  string `json:"idorg"`
}

// Organization - entity for organization
type Organization struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Status      string   `json:"status"`
	IDsCategory []string `json:"categories"`
}
```

### Resolvers
Resolver
```
package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"mini-seller/application/graph/model"
	"mini-seller/domain/packages/catalogpkg"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CatalogUseCase catalogpkg.IUseCase

	organizations []*model.Organization
	product       *model.Product
}
```

Main resolver for query
```
package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/model"

	"github.com/prometheus/common/log"
)

func (r *queryResolver) Organizations(ctx context.Context) ([]*model.Organization, error) {
	orgs, err := r.CatalogUseCase.GetOrganizationList(ctx)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if len(orgs) == 0 {
		return nil, nil
	}

	organizations := make([]*model.Organization, len(orgs))
	for i, org := range orgs {
		organizations[i] = &model.Organization{
			ID:          org.ID,
			Name:        org.Name,
			Email:       org.Email,
			Phone:       org.Phone,
			Status:      org.Status,
			IDsCategory: org.IDsCategory,
		}
	}

	return organizations, err
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	prod, err := r.CatalogUseCase.GetProductDetail(ctx, id)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if prod == nil {
		return nil, nil
	}

	product := &model.Product{
		ID:             prod.ID,
		IDCategory:     prod.IDCategory,
		IDOrganization: prod.IDOrganization,
		Name:           prod.Name,
		Price:          prod.Price,
		Count:          prod.Count,
		Status:         prod.Status,
	}

	return product, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
```

Resolver for entities
```
package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/model"

	"github.com/prometheus/common/log"
)

func (r *categoryResolver) Products(ctx context.Context, obj *model.Category) ([]*model.Product, error) {
	IDOrg := ""
	if obj.IDOrg != "" {
		IDOrg = obj.IDOrg
	}

	prodList, err := r.CatalogUseCase.GetProductList(ctx, []string{IDOrg}, []string{obj.ID})
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if prodList == nil {
		return nil, nil
	}

	products := make([]*model.Product, len(prodList))
	for i, product := range prodList {
		products[i] = &model.Product{
			ID:             product.ID,
			IDCategory:     product.IDCategory,
			IDOrganization: product.IDOrganization,
			Name:           product.Name,
			Price:          product.Price,
			Count:          product.Count,
			Status:         product.Status,
		}
	}

	return products, nil
}

func (r *organizationResolver) Categories(ctx context.Context, obj *model.Organization) ([]*model.Category, error) {
	catList, err := r.CatalogUseCase.GetCategoryList(ctx, obj.IDsCategory)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	if catList == nil {
		return nil, nil
	}

	categories := make([]*model.Category, len(catList))
	for i, category := range catList {
		categories[i] = &model.Category{
			ID:     category.ID,
			Name:   category.Name,
			Status: category.Status,
			IDOrg:  obj.ID,
		}
	}

	return categories, nil
}

func (r *organizationResolver) Products(ctx context.Context, obj *model.Organization) ([]*model.Product, error) {
	return ctxLoaders(ctx).productsByOrganization.Load(obj.ID)
}

func (r *productResolver) Category(ctx context.Context, obj *model.Product) (*model.Category, error) {
	cat, err := r.CatalogUseCase.GetCategoryDetail(ctx, obj.IDCategory)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, nil
	}

	category := &model.Category{
		ID:     cat.ID,
		Name:   cat.Name,
		Status: cat.Status,
		IDOrg:  obj.IDOrganization,
	}

	return category, nil
}

func (r *productResolver) Organization(ctx context.Context, obj *model.Product) (*model.Organization, error) {
	org, err := r.CatalogUseCase.GetOrganizationDetail(ctx, obj.IDOrganization)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, nil
	}

	organization := &model.Organization{
		ID:          org.ID,
		Name:        org.Name,
		Email:       org.Email,
		Phone:       org.Phone,
		Status:      org.Status,
		IDsCategory: org.IDsCategory,
	}

	return organization, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type categoryResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
```

### Web server
Launching of web api:
- Reading of configuration
- Connection to database
- Initialization of repositories
- Initialization of use cases
- Initialization of web framework
- Launching of web server

```
package main

import (
	"log"
	"mini-seller/application/graph/generated"
	"mini-seller/application/graph/resolvers"
	"mini-seller/domain/packages/catalogpkg/catalogrepository"
	"mini-seller/domain/packages/catalogpkg/catalogusecase"
	"mini-seller/infrastructure/mongohelper"
	"mini-seller/infrastructure/viperhelper"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

func main() {
	vip := viperhelper.Viper{ConfigType: "", ConfigName: "", ConfigPath: "infrastructure/viperhelper"}
	vip.Read()

	db, err := mongohelper.Connect("")
	if err != nil {
		log.Fatal(err)
	}

	catalogRepository := catalogrepository.NewCatalogRepository(db)
	catalogUseCase := catalogusecase.NewCatalogUseCase(catalogRepository)

	resolver := resolvers.Resolver{CatalogUseCase: catalogUseCase}

	router := http.NewServeMux()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver})))

	port := viper.GetString("WEB_PORT")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, resolvers.LoaderMiddleware(catalogUseCase, router)))
}
```

Launch server
```
go run application/server.go
```

### GraphQL
GraphQL is kind of web api, it provide a query language, understandable documentation for your requests and has pretty good playground.

With GraphQL you could select fields that you need from backend and you could create resolvers for related entities. Example for product:

```
query product($id: String!) {
    product(id: $id) {
        id
        name
        price
        count
        category{
            id
            name
        }
        organization {
            id
            name
        }
    }
}
```

Graphql have powerful query language, you could use Aliases and get different data in one request and you could use Fragment for common fields.
Query
```
query products($id_cola: String! $id_salad: String!) {
    cola: product(id: $id_cola) {
        ...detailFields
    }
  	salad: product(id: $id_salad) {
        ...detailFields
    }
}

fragment detailFields on Product {
    id
    name
    price
    count
    category{
        id
        name
    }
    organization {
        id
        name
    }
}
```
Variables
```
{
    "id_cola": "604497558ffcad558eb8e1f5",
    "id_salad": "604497558ffcad558eb8e1f4"
}
```

Result
```
{
    "data": {
        "cola": {
            "id": "604497558ffcad558eb8e1f5",
            "name": "Cola",
            "price": 100,
            "count": 100,
            "category": {
                "id": "604488100f719d9c76a28fe3",
                "name": "Drinks"
            },
            "organization": {
                "id": "6043d76e94df8de741c2c0d6",
                "name": "restaurant"
            }
        },
        "salad": {
            "id": "604497558ffcad558eb8e1f4",
            "name": "Salad Cesar",
            "price": 200,
            "count": 10,
            "category": {
                "id": "604488100f719d9c76a28fe6",
                "name": "Salad"
            },
            "organization": {
                "id": "6043d76e94df8de741c2c0d6",
                "name": "restaurant"
            }
        }
    }
}
```





### Validation
validator.v9 - package very helpful for validation of structures
```
go get gopkg.in/go-playground/validator.v9
```