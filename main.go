package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// main vars/consts should not be exported
// if exported then import cycle error will occur
var (
	server *http.Server
	router *gin.Engine
)

func main() {

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Println("server is shutdown ...")
		}
	}()

	//stopping the server
	stopch := make(chan bool)
	go handleShutDown(stopch)
	<-stopch

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalln("error while shutting down the server: ", err.Error())
	}

}
