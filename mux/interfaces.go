package mux

import (
	"net/http"

	"TravelAssistant/requesthandlers/alexa"
)

type RequestHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, rh alexa.RequestHandler)
}
