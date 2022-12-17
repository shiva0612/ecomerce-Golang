package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id" json:"-"`
	OrderID       string             `bson:"order_id" json:"order_id"`
	UserID        string             `bson:"user_id" `
	AddressID     string             `bson:"address_id" json:"address_id"`
	PaymentOption string             `bson:"payment_option" json:"payment_option"`
	Email         string             `bson:"email" json:"email"`
	Phone         string             `bson:"phone" json:"phone"`
	TotalSum      string             `bson:"total_sum" json:"total_sum"`
	Date          time.Time          `json:"date" bson:"date"`
	//embeded
	CartProducts []CartProduct `bson:"cart_products" json:"cart_products"`
}

type CheckoutCart struct {
	AddressID     string ` json:"address_id"`
	PaymentOption string ` json:"payment_option"`
	Email         string ` json:"email"`
	Phone         string ` json:"phone"`
}
