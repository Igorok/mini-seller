// Package orderentity - entity for order
package orderentity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderMongo - model for mongo
type OrderMongo struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	IDCustomer       primitive.ObjectID `bson:"_id_cust"`
	IDOrganization   primitive.ObjectID `bson:"_id_ord"`
	Basket           []BasketItemMongo
	Type             string
	Status           string
	Price            float64
	DateCreate       time.Time
	Comment          string
	DateCustomerWish time.Time
	Address          string
}

// BasketItemMongo - model for mongo
type BasketItemMongo struct {
	ID    primitive.ObjectID `bson:"_id"`
	Price float64
	Count int
}

// ToOrder - convert model to entity
func ToOrder(om *OrderMongo) *Order {
	ID := ""
	if !om.ID.IsZero() {
		ID = om.ID.Hex()
	}

	IDCustomer := om.IDCustomer.Hex()
	IDOrganization := om.IDOrganization.Hex()

	Basket := make([]BasketItem, len(om.Basket))
	for i, product := range om.Basket {
		Basket[i] = *ToBasketItem(&product)
	}

	return &Order{
		ID:               ID,
		IDCustomer:       IDCustomer,
		IDOrganization:   IDOrganization,
		Basket:           Basket,
		Type:             om.Type,
		Status:           om.Status,
		Price:            om.Price,
		DateCreate:       om.DateCreate,
		Comment:          om.Comment,
		DateCustomerWish: om.DateCustomerWish,
		Address:          om.Address,
	}
}

// ToOrderMongo - convert entity to model
func ToOrderMongo(o *Order) *OrderMongo {
	var ID primitive.ObjectID
	if o.ID != "" {
		ID, _ = primitive.ObjectIDFromHex(o.ID)
	}

	IDCustomer, _ := primitive.ObjectIDFromHex(o.IDCustomer)
	IDOrganization, _ := primitive.ObjectIDFromHex(o.IDOrganization)

	BasketMongo := make([]BasketItemMongo, len(o.Basket))
	for i, product := range o.Basket {
		BasketMongo[i] = *ToBasketItemMongo(&product)
	}

	return &OrderMongo{
		ID:               ID,
		IDCustomer:       IDCustomer,
		IDOrganization:   IDOrganization,
		Basket:           BasketMongo,
		Type:             o.Type,
		Status:           o.Status,
		Price:            o.Price,
		DateCreate:       o.DateCreate,
		Comment:          o.Comment,
		DateCustomerWish: o.DateCustomerWish,
		Address:          o.Address,
	}
}

// ToBasketItem - convert model to entity
func ToBasketItem(bim *BasketItemMongo) *BasketItem {
	return &BasketItem{
		ID:    bim.ID.Hex(),
		Price: bim.Price,
		Count: bim.Count,
	}
}

// ToBasketItemMongo - convert entity to model
func ToBasketItemMongo(bi *BasketItem) *BasketItemMongo {
	ID, _ := primitive.ObjectIDFromHex(bi.ID)

	return &BasketItemMongo{
		ID:    ID,
		Price: bi.Price,
		Count: bi.Count,
	}
}
