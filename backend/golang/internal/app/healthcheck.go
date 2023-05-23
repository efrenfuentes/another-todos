package app

import (
	"net/http"
)

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Env,
			"version":     version,
		},
	}

	err := app.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}
