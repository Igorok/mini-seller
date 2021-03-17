// Package productentity - entities for product
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

// ProductForList - entity for catalog of products
type ProductForList struct {
	ID           string
	Category     CategoryForProduct
	Organization OrganizationForProduct
	Name         string
	Price        int
	Count        int
	Status       string
}

// CategoryForProduct - for catalog of products
type CategoryForProduct struct {
	ID   string
	Name string
}

// OrganizationForProduct - for catalog of products
type OrganizationForProduct struct {
	ID    string
	Name  string
	Phone string
	Email string
}

// ProductList - entity with data for catalog
type ProductList struct {
	Products []*ProductForList
	Count    int64
}
