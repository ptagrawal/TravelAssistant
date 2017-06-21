package travelAssistant

import "TravelAssistant/protocol/alexa"

type requestHandler interface {
	handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
