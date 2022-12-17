package database

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoURI string
	client   *mongo.Client

	userColl     *mongo.Collection
	addColl      *mongo.Collection
	productsColl *mongo.Collection
	orderCOll    *mongo.Collection
)
