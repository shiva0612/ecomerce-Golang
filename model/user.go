package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	UserType  string             `json:"user_type"  bson:"user_type" `
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Phone     string             `json:"phone" bson:"phone"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Orders    []string           `json:"orders" bson:"orders"`
	//TODO:this will also create _id field automatically, which is to be ignored while projecting
	CartProducts []CartProduct `bson:"cart_products"`
}

type CartProduct struct {
	ProductID  string  `bson:"product_id" json:"product_id"`
	Quantity   int     `bson:"quantity" json:"quantity"`
	ProductSum float64 `bson:"product_sum" json:"product_sum"`
}

type Address struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	AddressID string             `bson:"address_id" json:"address_id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	Hno       string             `json:"hno" bson:"hno"`
	City      string             `json:"city" bson:"city"`
	State     string             `json:"state" bson:"state"`
	Pincode   string             `json:"pincode" bson:"pincode"`
}

type User_Signup struct {
	UserType string `form:"user_type" json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Name     string `form:"name" json:"name" validate:"required,min=1,max=50"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Phone    string `form:"phone" json:"phone" validate:"required,max=10"`
	Username string `form:"username" json:"username" validate:"required,max=50"`
	Password string `form:"password" json:"password" validate:"required,min=10,max=20"`
}

type User_Login struct {
	UserType string `form:"user_type" json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Username string `form:"username" json:"username" validate:"required,max=50"`
	Password string `form:"password" json:"password" validate:"required,min=10,max=20"`
}

// keep only those information in claims which are required to access other webpages
// dont keep to many fields bcz,
// tokens are parsed to get claims for every authorized webpage
// if userclaims are full of fields they will take up mry
type UserClaims struct {
	UserID string
	// Username string
	// UserType string
	// Email    string
	// Phone    string
	jwt.StandardClaims
}
