package models

import "time"

type DefaultResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Warning   string `json:"warnings,omitempty"`
}

func CreateResponseMessage(message string) DefaultResponse {
	return DefaultResponse{Message: message, Timestamp: time.Now().String()}
}
