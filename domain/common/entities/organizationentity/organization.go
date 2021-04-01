// Package organizationentity - entities for organization
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
