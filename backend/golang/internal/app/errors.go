package app

import (
	"fmt"
	"net/http"
)

// The logError() method is a generic helper for logging an error message. Later
// we'll upgrade this to use structured logging, and record additional information
// about the request including the HTTP method and URL.
func (app *Application) LogError(r *http.Request, err error) {
	app.Logger.Println(err)
}

// The errorResponse() method is a generic helper for sending JSON-formatted error
// messages to the client with a given status code. Note that we're using an interface{}
// type for the message parameter, rather than just a string type, as this gives us
// more flexibility over the values that we can include in the response.
func (app *Application) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := Envelope{"error": message}

	// Write the response using the writeJSON() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := app.WriteJSON(w, status, env, nil)
	if err != nil {
		app.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (app *Application) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.LogError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send 404 Not Found status code and
// JSON response to the client.
func (app *Application) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the request resource could not be found"
	app.ErrorResponse(w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to client
func (app *Application) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// The FailedValidationResponse() method will be used to send a 422 Unprocessable Entity
func (app *Application) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
