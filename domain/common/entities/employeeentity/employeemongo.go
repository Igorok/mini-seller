// Package employeeentity - package for employee
package employeeentity

import "go.mongodb.org/mongo-driver/bson/primitive"

// EmployeeMongo - model for database
type EmployeeMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string
	Sname     string
	Email     string
	Password  string
	Phone     string
	Status    string
	Positions []PositionMongo
}

// PositionMongo - model for database
type PositionMongo struct {
	IDOrganization primitive.ObjectID `bson:"_id_org"`
	Roles          []string
}

// ToEmployee - convert model to entity
func ToEmployee(em *EmployeeMongo) *Employee {
	ID := ""
	if !em.ID.IsZero() {
		ID = em.ID.Hex()
	}

	Positions := make([]Position, len(em.Positions))
	for i, position := range em.Positions {
		Positions[i] = *ToPosition(&position)
	}

	return &Employee{
		ID:        ID,
		Name:      em.Name,
		Sname:     em.Sname,
		Email:     em.Email,
		Password:  em.Password,
		Phone:     em.Phone,
		Status:    em.Status,
		Positions: Positions,
	}
}

// ToEmployeeMongo - convert model to entity
func ToEmployeeMongo(e *Employee) *EmployeeMongo {
	var ID primitive.ObjectID
	if e.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(e.ID)
	}

	Positions := make([]PositionMongo, len(e.Positions))
	for i, position := range e.Positions {
		Positions[i] = *ToPositionMongo(&position)
	}

	return &EmployeeMongo{
		ID:        ID,
		Name:      e.Name,
		Sname:     e.Sname,
		Email:     e.Email,
		Password:  e.Password,
		Phone:     e.Phone,
		Status:    e.Status,
		Positions: Positions,
	}
}

// ToPosition - convert model to entity
func ToPosition(pm *PositionMongo) *Position {
	return &Position{
		IDOrganization: pm.IDOrganization.Hex(),
		Roles:          pm.Roles,
	}
}

// ToPositionMongo - convert entity to model
func ToPositionMongo(p *Position) *PositionMongo {
	IDOrganization, _ := primitive.ObjectIDFromHex(p.IDOrganization)

	return &PositionMongo{
		IDOrganization: IDOrganization,
		Roles:          p.Roles,
	}
}
