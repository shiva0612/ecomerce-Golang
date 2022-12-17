package main

import (
	"log"
	"os"
	"os/signal"
	"shiva/database"
	"syscall"
)

func handleShutDown(stop chan bool) {
	funcName := "handleShutDown"
	exitch := make(chan os.Signal, 10)
	signal.Notify(exitch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGHUP, syscall.SIGKILL, os.Interrupt, syscall.SIGTERM)
	<-exitch

	preShutdownTasks()

	log.Printf("[%s] shutting down server", funcName)
	stop <- true
}

func preShutdownTasks() {
	funcName := "preShutdownTasks"
	log.Printf("[%s] performing pre-shutdown tasks", funcName)

	database.DisconnectDB()

}
