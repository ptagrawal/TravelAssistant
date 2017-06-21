package alexa

import "TravelAssistant/protocol/alexa"

type RequestHandler interface {
	HandleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error)
}
