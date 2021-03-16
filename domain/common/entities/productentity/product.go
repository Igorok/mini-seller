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

// ProductForCatalog - entity for catalog of products
type ProductForCatalog struct {
	ID           string
	Category     CategoryForCatalog
	Organization OrganizationForCatalog
	Name         string
	Price        int
	Count        int
	Status       string
}

// CategoryForCatalog - for catalog of products
type CategoryForCatalog struct {
	ID   string
	Name string
}

// OrganizationForCatalog - for catalog of products
type OrganizationForCatalog struct {
	ID    string
	Name  string
	Phone string
	Email string
}

// CatalogData - entity with data for catalog
type CatalogData struct {
	Products []*ProductForCatalog
	Count    int64
}
