// Package employeeentity - package for employee
package employeeentity

// Employee - entity
type Employee struct {
	ID        string     `validate:"omitempty"`
	Name      string     `validate:"required"`
	Sname     string     `validate:"required"`
	Email     string     `validate:"email"`
	Password  string     `validate:"email"`
	Phone     string     `validate:"required"`
	Status    string     `validate:"required"`
	Positions []Position `validate:"required"`
}

// Position - entity for roles in organization
type Position struct {
	IDOrganization string   `validate:"required"`
	Roles          []string `validate:"required"`
}
