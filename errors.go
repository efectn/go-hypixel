package gohypixel

import (
	"encoding/json"
	"net/http"

	"github.com/efectn/go-hypixel/models"
)

// A struct to handle error responses.
type Error struct {
	Code    int            `json:"code"`    // Status code of error.
	Message string         `json:"message"` // Error message of error.
	Data    map[string]any `json:"data"`    // Additional datas for error.
}

// Error makes it compatible with the `error` interface.
func (e Error) Error() string {
	return e.Message
}

// NewError is a function to create error instance.
func NewError(code int, message string, data ...map[string]any) *Error {
	err := &Error{
		Code:    code,
		Message: message,
	}

	if len(data) > 0 {
		err.Data = data[0]
	}

	return err
}

// checkHTTPErrors simplfies error processing of response.
func checkHTTPErrors(resp *http.Response, body string) error {
	if resp.StatusCode != http.StatusOK {
		var jsonOutput models.Error
		if err := json.Unmarshal([]byte(body), &jsonOutput); err != nil {
			return err
		}

		return NewError(resp.StatusCode, jsonOutput.Cause)
	}

	return nil
}

// Errors for responses
var ErrInvalidAPI = NewError(403, "Invalid API key.")
var ErrThrottle = NewError(429, "Key throttle.")
var ErrMissingFields = NewError(400, "Missing one or more fields.")
var ErrMalformedUUID = NewError(422, "Malformed UUID.")
var ErrNoResultFound = NewError(404, "No result was found.")
var ErrPageNotFound = NewError(404, "Page not found.")
var ErrLeaderboardNotAvailable = NewError(503, "Leaderboard data has not yet been populated.")
var ErrNoBingoFound = NewError(404, "No bingo data could be found.")
