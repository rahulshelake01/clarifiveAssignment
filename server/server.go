package server

import (
	"clarifiveassignment/api/v1/repository"
	"clarifiveassignment/api/v1/router"
	"clarifiveassignment/api/v1/services"
	"clarifiveassignment/config"
	"clarifiveassignment/database"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type ServerInterface interface {
	Start()
}

type serverStrcut struct{}

func Server() ServerInterface {
	return serverStrcut{}
}

func (s serverStrcut) Start() {

	dbConfig := config.GetDBConfig()

	Connection := database.DBConnection(dbConfig)
	db, err := Connection.InitDB()

	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	familyRepository := repository.FamilyRepository(db)
	familyService := services.FamilyService(familyRepository)
	router := router.RouteHandler(familyService)

	log.Println("Social media server srated on http://" + os.Getenv("SOCIAL_MEDIA_HOST") + ":" + os.Getenv("SOCIAL_MEDIA_PORT"))

	if err := http.ListenAndServe(os.Getenv("SOCIAL_MEDIA_HOST")+":"+os.Getenv("SOCIAL_MEDIA_PORT"), router); err != nil {
		log.Fatal("Failed to start social-media server...")
	}
}
