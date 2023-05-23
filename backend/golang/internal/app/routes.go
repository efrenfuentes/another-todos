package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	// Convert the notFoundResponse() helper to http.Handler using the
	// http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	// Not Found responses.
	router.NotFound = http.HandlerFunc(app.NotFoundResponse)

	// Automatic OPTIONS responses and CORS
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "Content-Type")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	// Likewise, convert the methodNotAllowedResponse() helper to http.Handler and set
	// it as the custom error handler for 405 Method Not Allowed responses.
	router.MethodNotAllowed = http.HandlerFunc(app.MethodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/todos", app.ListTodosHandler)
	router.HandlerFunc(http.MethodPost, "/v1/todos", app.CreateTodoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/todos/:id", app.ShowTodoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/todos/:id", app.UpdateTodoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/todos/:id", app.DeleteTodoHandler)

	return app.enableCORS(app.logRequest(router))
}
