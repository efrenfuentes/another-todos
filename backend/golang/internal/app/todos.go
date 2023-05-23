package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/efrenfuentes/todo-backend-go/internal/data"
	"github.com/efrenfuentes/todo-backend-go/internal/validator"
)

func (app *Application) ListTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := app.Models.Todos.GetAll()
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	err = app.WriteJSON(w, http.StatusOK, Envelope{"todos": todos}, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}

func (app *Application) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
		Order     int    `json:"order"`
		Url       string `json:"url"`
	}

	err := app.ReadJSON(w, r, &input)
	if err != nil {
		app.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	todo := &data.Todo{
		Title:     input.Title,
		Completed: input.Completed,
		Order:     input.Order,
		Url:       input.Url,
	}

	v := validator.New()

	if data.ValidateTodo(v, todo); !v.Valid() {
		app.FailedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.Models.Todos.Insert(todo)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	// When sending a HTTP response, we want to include a Location header to
	// let the client known which URL they can find the newly created resource.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/todos/%d", todo.ID))

	// Write a JSON response with a 201 Created status code and the Location header.
	err = app.WriteJSON(w, http.StatusCreated, Envelope{"todo": todo}, headers)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}

func (app *Application) ShowTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadIDParam(r)
	if err != nil {
		app.NotFoundResponse(w, r)
		return
	}

	todo, err := app.Models.Todos.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.NotFoundResponse(w, r)
		default:
			app.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = app.WriteJSON(w, http.StatusOK, Envelope{"todo": todo}, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}

func (app *Application) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadIDParam(r)
	if err != nil {
		app.NotFoundResponse(w, r)
		return
	}

	var input struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
		Order     *int    `json:"order"`
		Url       *string `json:"url"`
	}

	err = app.ReadJSON(w, r, &input)
	if err != nil {
		app.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := app.Models.Todos.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.NotFoundResponse(w, r)
		default:
			app.ServerErrorResponse(w, r, err)
		}
		return
	}

	// Copy the data from the input struct to the todo struct.
	if input.Title != nil {
		todo.Title = *input.Title
	}
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}
	if input.Order != nil {
		todo.Order = *input.Order
	}
	if input.Url != nil {
		todo.Url = *input.Url
	}

	v := validator.New()

	if data.ValidateTodo(v, todo); !v.Valid() {
		app.FailedValidationResponse(w, r, v.Errors)
		return
	}

	// Update the todo record in the database.
	err = app.Models.Todos.Update(todo)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	// Write the JSON response with a 200 OK status code.
	err = app.WriteJSON(w, http.StatusOK, Envelope{"todo": todo}, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}

func (app *Application) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadIDParam(r)
	if err != nil {
		app.NotFoundResponse(w, r)
		return
	}

	err = app.Models.Todos.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.NotFoundResponse(w, r)
		default:
			app.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = app.WriteJSON(w, http.StatusOK, Envelope{"message": "todo successfully deleted"}, nil)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
	}
}
