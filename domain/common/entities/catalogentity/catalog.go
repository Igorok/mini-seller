// Package catalogentity - entities for catalog
package catalogentity

// ProductInfo - entity for product
type ProductInfo struct {
	ID             string
	IDCategory     string
	IDOrganization string
	Name           string
	Price          int
	Count          int
	Status         string
}

// CategoryInfo - entity for category of products
type CategoryInfo struct {
	ID     string
	Name   string
	Status string
}

// OrganizationInfo - entity for organization
type OrganizationInfo struct {
	ID          string
	Name        string
	Email       string
	Phone       string
	Status      string
	IDsCategory []string
}
