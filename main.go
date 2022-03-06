package main

import (
	"clarifiveassignment/config"
	"clarifiveassignment/server"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}
	config.LoadConfig()

}

func initDB() {

}

func main() {
	server := server.Server()
	server.Start()
}
