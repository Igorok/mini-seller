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
	Category       CategoryInfo
	Organization   OrganizationInfo
}

// CategoryInfo - entity for category of products
type CategoryInfo struct {
	ID       string
	Name     string
	Status   string
	Products []ProductInfo
}

// OrganizationInfo - entity for organization
type OrganizationInfo struct {
	ID         string
	Name       string
	Email      string
	Phone      string
	Status     string
	Products   []ProductInfo
	Categories CategoryInfo
}
