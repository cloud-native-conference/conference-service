package restapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ListenAndServe(listeningPort string, controller *Controller) error {
	router := httprouter.New()
	RegisterHandlers(router, controller)
	return http.ListenAndServe(":"+listeningPort, router)
}

func RegisterHandlers(router *httprouter.Router, controller *Controller) {
	router.GET("/", controller.getConferences)
	router.GET("/:uniqueName", controller.getConference)
	router.POST("/", controller.createConference)
}
