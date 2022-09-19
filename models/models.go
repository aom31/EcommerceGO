package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json: "_id" bson: "_id"`
	FirstName     *string            `json: "firstname" validate: "required, min=2, max=15"`
	LastName      *string            `json: "lastname" validate: "required, min=2, max=15"`
	Password      *string            `json: "password" validate: "required, min=6"`
	Email         *string            `json: "email" validate: "email , required"`
	Telephone     int                `json: "telephone" validate : "required"`
	Token         *string            `json: "token"`
	Refresh_token *string            `json: "refesh_token"`
	CreateAt      time.Time          `json: "create_at"`
	UpdateAt      time.Time          `json: "update_at"`
	UserID        string             `json: "user_id"`
	UserCart      []ProductUser      `json: "user_cart" bson: "user_cart"`
	Address       []AddressDetail    `json: "address" bson: "address"`
	OrderStatus   []Order            `json: "order_status" bson: "order_status"`
}

type Product struct {
	ProductID   primitive.ObjectID `json: "product_id" bson: "product_id"`
	ProductName *string            `json: "product_name"`
	Price       float64            `json: "price"`
	Rating      *uint8             `json: "rating"`
	Image       *string            `json: "image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson : "product_id"`
	ProductName *string            `json: "product_name"`
	Price       float64            `json: "price"`
	Rating      *uint              `json: "rating"`
	Image       *string            `json: "image"`
}

type AddressDetail struct {
	AddressID primitive.ObjectID `json: "_id" bson: "_id"`
	House     *string            `json: "house"`
	Street    *string            `json: "street"`
	City      *string            `json: "city"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson: "_id"`
	OrderCart     []ProductUser      `bson: "order_cart"`
	OrderAt       time.Time          `json: "order_at"`
	Price         int                `json: "price"`
	Discount      *int               `json: "discount"`
	PaymentMethod Payment            `json: "payment_method"`
}

type Payment struct {
	Digital bool `json: "digital"`
	COD     bool `json: "cod"`
}
