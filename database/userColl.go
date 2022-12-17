package database

import (
	"context"
	"errors"
	"log"
	"shiva/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user_signup *model.User_Signup) error {

	funcName := "CreateUser"
	user := model.User_from_userSingup(user_signup)
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()
	user.Orders = []string{}
	user.CartProducts = make([]model.CartProduct, 0)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err := userColl.InsertOne(context.TODO(), user)
	if err != nil {
		log.Printf("[%s]: error while inserting user in DB: %s", funcName, err.Error())
		return err
	}

	return nil
}

func GetUser(username string) (*model.User, error) {
	funcName := "GetUser"
	filter := bson.M{"username": username}
	res := userColl.FindOne(context.TODO(), filter)
	if res.Err() == mongo.ErrNoDocuments {
		log.Printf("[%s]: no user found", funcName)
		return nil, errors.New(ErrUSER_NOT_FOUND)
	}
	user := &model.User{}
	res.Decode(user)
	return user, res.Err()
}

func GetProfile(userID string) (*model.User, error) {

	_id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": _id}
	res := userColl.FindOne(context.Background(), filter)
	if res.Err() != nil {
		log.Println("error while getting profile", res.Err().Error())
		return nil, res.Err()
	}
	user := new(model.User)
	err := res.Decode(user)
	if err != nil {
		log.Println("error while decoding user: ", err.Error())
		return nil, err
	}

	return user, nil
}

func DeleteUser(userID string) error {

	//delete from user,address & orders
	filter := bson.M{"user_id": userID}
	_, err := orderCOll.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println("error deleting from order coll: ", err.Error())
		return err
	}
	_, err = addColl.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println("error deleting from address coll: ", err.Error())
		return err
	}
	_, err = userColl.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("error deleting from user coll: ", err.Error())
		return err
	}
	return nil
}

// user exists -> returns true,nil
//
// if err or user is not found -> return false,error
func CheckIfUserExists(username, email, phone string) (bool, error) {
	funcName := "CheckIfUserExists"

	filter := bson.M{
		"$or": bson.A{
			bson.M{"username": username},
			bson.M{"email": email},
			bson.M{"phone": phone},
		},
	}
	res := userColl.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Printf("[%s]: %s", funcName, res.Err().Error())
		return true, res.Err()
	}
	return true, errors.New(ErrUser_ALREADY_EXIST)

}

func AddAddress(userID string, address []model.Address) (err error) {
	pis := []interface{}{}
	for _, pi := range address {
		pi.ID = primitive.NewObjectID()
		pi.AddressID = pi.ID.Hex()
		pi.UserID = userID
		pis = append(pis, pi)
	}

	_, err = addColl.InsertMany(context.Background(), pis)
	if err != nil {
		log.Println("error adding address: ", err.Error())
		return
	}
	return
}

func GetAddress(userID string) (address []model.Address, err error) {
	address = []model.Address{}
	filter := bson.M{
		"user_id": userID,
	}
	cur, err := addColl.Find(context.Background(), filter)
	if err != nil {
		log.Println("error getting address: ", err.Error())
		return
	}
	err = cur.All(context.Background(), &address)
	if err != nil {
		log.Println("error while unmarshalling address: ", err.Error())
		return
	}
	return
}
