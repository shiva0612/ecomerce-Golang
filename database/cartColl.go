package database

import (
	"context"
	"log"
	"shiva/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddProductToCart(userID, productID string, quantity int) (err error) {

	cartProduct := model.CartProduct{
		ProductID: productID,
		Quantity:  quantity,
	}

	_id, _ := primitive.ObjectIDFromHex(userID)
	update := bson.M{
		"$push": bson.M{
			"cart_products": cartProduct,
		},
	}
	_, err = userColl.UpdateByID(context.Background(), _id, update)
	if err != nil {
		log.Println("error while adding product to cart: ", err.Error())
		return
	}

	return
}

func RemoveProductFromCart(userID, productID string) (err error) {

	_id, _ := primitive.ObjectIDFromHex(userID)
	update := bson.M{
		"$pull": bson.M{
			"cart_products": bson.M{
				"product_id": productID,
			},
		},
	}
	_, err = userColl.UpdateOne(context.Background(), _id, update)
	if err != nil {
		log.Println("error while removing product from cart: ", err.Error())
		return
	}

	return
}

func ViewCart(userID string) ([]model.CartProduct, error) {
	_id, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{
		"_id": _id,
	}
	findOpt := &options.FindOneOptions{
		Projection: bson.M{
			"cart_products": 1,
			"_id":           0,
		},
	}

	res := userColl.FindOne(context.Background(), filter, findOpt)
	if res.Err() != nil {
		log.Println("error while viewing cart: ", res.Err().Error())
		return nil, res.Err()
	}
	// check := bson.M{}
	// res.Decode(&check)
	// b, _ := json.MarshalIndent(check, " ", "  ")
	// log.Println(string(b))
	// log.Println()
	// return nil, nil
	//TODO:check if this is the only way or we can do something with aggregation pipeline
	cartProucts := make([]model.CartProduct, 0)
	result := map[string][]model.CartProduct{}
	err := res.Decode(&result)
	if err != nil {
		log.Println("error while decoding cartItems: ", err.Error())
		return nil, err
	}
	cartProucts = result["cart_products"]
	return cartProucts, nil

}

func CheckoutCart(userID string) (address []model.Address, err error) {

	address = make([]model.Address, 0)

	filter := bson.M{
		"user_id": userID,
	}

	cur, err := addColl.Find(context.Background(), filter)
	if err != nil {
		log.Println("error while fetching address: ", err.Error())
		return
	}
	err = cur.All(context.Background(), &address)
	if err != nil {
		log.Println("error while decoding address: ", err.Error())
		return
	}

	return

}
