package main

import (
	"fmt"
	"log"
	"net/http"
	"shiva/database"
	"shiva/middleware"
	"shiva/routes"

	"github.com/gin-gonic/gin"

	"shiva/config"
)

func init() {

	err := config.LoadConfig("./config/jwtConfig.json")
	if err != nil {
		log.Fatalln("error loading the config file: ", err.Error())
	}

	getMongoURI()
	database.ConnectToDB()

	router = gin.New()

	routes.AuthRoutes(router)

	router.Use(middleware.Authenticate)
	routes.UserRoutes(router)
	routes.EcomRoutes(router)

	server = &http.Server{
		Addr:    ":" + config.Prjconfig.Port,
		Handler: router,
	}
}

func getMongoURI() {

	database.MongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
		config.Prjconfig.Mongo.User,
		config.Prjconfig.Mongo.Psw,
		config.Prjconfig.Mongo.Server,
		config.Prjconfig.Mongo.Port,
		config.Prjconfig.Mongo.Auth)

	if database.MongoURI == "" {
		log.Fatalln("empty database connection string URI")
		return
	}
	log.Println("mongodb URI = ", database.MongoURI)
}
