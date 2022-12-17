package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DisconnectDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalln("error while disconnect from databse" + err.Error())
	}
}

func ConnectToDB() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Fatalf("error getting database connection: %s", err.Error())
	}

	//get all collections objects
	userColl = client.Database(cDB).Collection(cUSERCOLL)
	addColl = client.Database(cDB).Collection(cADDRESSCOLL)
	productsColl = client.Database(cDB).Collection(cPRODUCTCOLL)
	orderCOll = client.Database(cDB).Collection(cORDERCOLL)

}
