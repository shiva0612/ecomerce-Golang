package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	ProductID string             `bson:"product_id" json:"product_id"`
	Name      string             `bson:"name" json:"name"`
	Category  string             `bson:"category" json:"category"`
	Price     string             `bson:"price" json:"price"`
	Image     string             `bson:"image" json:"image"`
	InStock   string             `bson:"in_stock" json:"in_stock"`
}
