package database

import (
	"context"
	"log"
	"shiva/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProducts(products []model.Product) (err error) {
	pis := []interface{}{}
	for _, pi := range products {
		pi.ID = primitive.NewObjectID()
		pi.ProductID = pi.ID.Hex()
		pis = append(pis, pi)
	}
	_, err = productsColl.InsertMany(context.Background(), pis)
	if err != nil {
		log.Println("error adding products: ", err.Error())
		return
	}
	return
}

func ProductsList(query map[string]string) (products []model.Product, err error) {
	products = make([]model.Product, 0)
	filter := bson.D{}

	for k, v := range query {
		e := bson.E{
			Key:   k,
			Value: v,
		}
		filter = append(filter, e)
	}

	cur, err := productsColl.Find(context.Background(), filter)
	if err != nil {
		log.Println("error while getting products: ", err.Error())
		return
	}
	err = cur.All(context.Background(), &products)
	if err != nil {
		log.Println("error while decoding the products: ", err.Error())
		return
	}

	return

}
