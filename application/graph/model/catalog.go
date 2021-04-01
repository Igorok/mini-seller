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
