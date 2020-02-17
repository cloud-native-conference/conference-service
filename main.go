package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloud-native-conference/conference-service/restapi"
	"github.com/cloud-native-conference/conference-service/service"
	"github.com/cloud-native-conference/conference-service/service/data"
)

const (
	MONGO_SERVER_URL_ENV_VAR_NAME = "MONGO_SERVER_URL"
	LISTENING_PORT                = "8080"
)

var MONGO_SERVER_URL = "mongodb://localhost:27017"

func main() {

	mongoServerUrlFromEnvironment := os.Getenv(MONGO_SERVER_URL_ENV_VAR_NAME)

	if mongoServerUrlFromEnvironment != "" {
		MONGO_SERVER_URL = mongoServerUrlFromEnvironment
	}

	databaseClient, err := data.New(MONGO_SERVER_URL)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	conferenceService := &service.ConferenceService{DatabaseClient: databaseClient}

	restApi := &restapi.Controller{ConferenceService: conferenceService}

	log.Fatal(restapi.ListenAndServe(LISTENING_PORT, restApi))
}
