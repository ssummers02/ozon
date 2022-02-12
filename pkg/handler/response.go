package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

// DefaultResponse is a JSON response in case of success.
type DefaultResponse struct {
	IsOK bool        `json:"is_ok"`
	Data interface{} `json:"data"`
}

// DefaultError is a JSON response in case of failure.
type DefaultError struct {
	Text string `json:"text"`
}

// SendErr sends a response to the client in case of success.
func SendErr(w http.ResponseWriter, code int, text string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(
		DefaultResponse{
			IsOK: false,
			Data: DefaultError{
				Text: text,
			},
		},
	)
}

// SendOK sends a response to the client in case of success.
func SendOK(w http.ResponseWriter, code int, p interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	// These two do not allow body
	_ = json.NewEncoder(w).Encode(
		DefaultResponse{
			IsOK: true,
			Data: p,
		},
	)
}
