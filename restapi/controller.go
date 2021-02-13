package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cloud-native-conference/conference-service/service"
	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	ConferenceService *service.ConferenceService
}

func (app *Controller) getConferences(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	conferences, err := app.ConferenceService.GetConferences()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(conferences); err != nil {
		encodeError := fmt.Errorf("Failed to encode conference object: %v, %w", conferences, err)
		http.Error(w, encodeError.Error(), http.StatusInternalServerError)
		return
	}

}

func (app *Controller) getConference(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uniqueName := params.ByName("uniqueName")
	conference, err := app.ConferenceService.GetConference(uniqueName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(conference); err != nil {
		encodeError := fmt.Errorf("Failed to encode conference object: %v, %w", conference, err)
		http.Error(w, encodeError.Error(), http.StatusInternalServerError)
		return
	}

}

func (app *Controller) createConference(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newConference service.Conference
	if err := json.NewDecoder(r.Body).Decode(&newConference); err != nil {
		decodeError := fmt.Errorf("Failed to decode conference body: %w", err)
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}
	conference, err := app.ConferenceService.CreateConference(&newConference)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(conference); err != nil {
		encodeErr := fmt.Errorf("Failed to encode conference object: %v, %w", conference, err)
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		return
	}
}
