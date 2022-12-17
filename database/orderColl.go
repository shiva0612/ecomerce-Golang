package database

import (
	"context"
	"log"
	"shiva/model"
	"time"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PlaceOrder(userID string, checkout model.CheckoutCart) (string, error) {
	newOrder := new(model.Order)

	newOrder.ID = primitive.NewObjectID()
	newOrder.OrderID = newOrder.ID.Hex()
	newOrder.UserID = userID
	newOrder.Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	copier.Copy(newOrder, checkout)

	err := copyCartItemsIntoOrder(userID, newOrder)
	if err != nil {
		log.Println("error in performing copyCartItemsIntoOrder in PlaceOrder: ", err.Error())
		return "", err
	}

	_, err = orderCOll.InsertOne(context.Background(), newOrder)
	if err != nil {
		log.Println("error creating order: ", err.Error())
		return "", err
	}

	emptyCartAndAddOrderID(userID, newOrder.OrderID)

	return newOrder.OrderID, nil

}

func copyCartItemsIntoOrder(userID string, newOrder *model.Order) error {
	_id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{
		"_id": _id,
	}
	projection := options.FindOneOptions{
		Projection: bson.M{
			"cart_products": 1,
			"_id":           0,
		},
	}

	res := userColl.FindOne(context.Background(), filter, &projection)
	if res.Err() != nil {
		log.Println("error while finding the user, when [copyCartItemsIntoOrder]: ", res.Err().Error())
		return res.Err()
	}
	result := map[string][]model.CartProduct{}
	err := res.Decode(&result)
	if err != nil {
		log.Println("error while decoding projection result of cart_products in [copyCartItemsIntoOrder]: ", err.Error())
		return err
	}
	newOrder.CartProducts = result["cart_products"]
	return nil
}

func emptyCartAndAddOrderID(userID, orderID string) {

	_id, _ := primitive.ObjectIDFromHex(userID)

	update := bson.M{
		"$set": bson.M{
			"cart_products": []model.CartProduct{},
		},
		"$push": bson.M{
			"orders": orderID,
		},
	}
	_, err := userColl.UpdateByID(context.Background(), _id, update)
	if err != nil {
		log.Println("error while emptying cart and updating orders in user: ", err.Error())
		return
	}

}

func ViewOrders(userID string) (orders []model.Order, err error) {

	orders = make([]model.Order, 0)

	filter := bson.M{
		"user_id": userID,
	}
	cur, err := orderCOll.Find(context.Background(), filter)
	if err != nil {
		log.Println("error while getting order: ", err.Error())
		return
	}

	err = cur.All(context.Background(), &orders)
	if err != nil {
		log.Println("error while decoding order: ", err.Error())
		return
	}

	return
}
