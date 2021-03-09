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

// InsertOrganizations - insert test data for organizations
func InsertOrganizations(db *mongo.Database) {
	type OrganizationData struct {
		Organizations []organizationentity.Organization
	}

	content, err := ioutil.ReadFile("infrastructure/mongohelper/testdata/organization.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var orgData OrganizationData
	err = json.Unmarshal(content, &orgData)
	if err != nil {
		log.Fatal("Unmarshal: ", err)
	}

	organizations := make([]interface{}, len(orgData.Organizations))
	for i, org := range orgData.Organizations {
		entity := organizationentity.ToMongo(&org)
		organizations[i] = entity
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Collection("organizations").DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Collection("organizations").InsertMany(ctx, organizations)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertProducts - insert test data for categories and products
func InsertProducts(db *mongo.Database) {
	type ProductsData struct {
		Categories []productcategoryentity.ProductCategory
		Products   []productentity.Product
	}

	content, err := ioutil.ReadFile("infrastructure/mongohelper/testdata/products.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var productData ProductsData
	err = json.Unmarshal(content, &productData)
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

// InsertEmployee - insert test data for employee
func InsertEmployee(db *mongo.Database) {
	type EmployeeData struct {
		Employees []employeeentity.Employee
	}

	content, err := ioutil.ReadFile("infrastructure/mongohelper/testdata/employee.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var emplData EmployeeData
	err = json.Unmarshal(content, &emplData)
	if err != nil {
		log.Fatal("Unmarshal: ", err)
	}

	employees := make([]interface{}, len(emplData.Employees))
	for i, empl := range emplData.Employees {
		entity := employeeentity.ToEmployeeMongo(&empl)
		employees[i] = entity
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Collection("employees").DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Collection("employees").InsertMany(ctx, employees)
	if err != nil {
		log.Fatal(err)
	}
}
