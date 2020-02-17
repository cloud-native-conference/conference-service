package service

import (
	"fmt"
	"log"

	"github.com/cloud-native-conference/conference-service/service/data"
)

type Conference struct {
	UniqueName  string `json:"uniqueName"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type ConferenceService struct {
	DatabaseClient data.DatabaseClient
}

func (sevice *ConferenceService) GetConference(uniqueName string) (*Conference, error) {
	log.Printf("Get Conference; uniqueId: %s", uniqueName)
	return nil, fmt.Errorf("Could not find conference with uniqueId: %s", uniqueName)
}

func (service *ConferenceService) UpdateConference(conference *Conference) (*Conference, error) {
	log.Printf("Update Conference; uniqueId: %s", conference.UniqueName)
	return nil, nil
}

func (service *ConferenceService) CreateConference(conference *Conference) (*Conference, error) {
	log.Printf("Create Conference; uniqueId: %s", conference.UniqueName)
	conferenceStorageModel := &data.ConferenceStorageModel{
		UniqueName:  conference.UniqueName,
		DisplayName: conference.DisplayName,
		Description: conference.Description,
	}
	err := service.DatabaseClient.CreateConference(conferenceStorageModel)
	if err != nil {
		return nil, fmt.Errorf("Could not create conference %v: %w", *conference, err)
	}
	return conference, nil
}
