package main

import (
	"company/internal/data"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Envelope map[string]interface{}

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//I need to decode the request body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorLogger.Println("Decoding request body", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//validate the password and email provided by the client
	// Lookup the user record based on the email address. If no matching user was
	// found, then we call the app.invalidCredentialsResponse() helper to send a 401
	// Unauthorized response to the client (we will create this helper in a moment).
	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.errorLogger.Println("Invalid email address", err)
			http.Error(w, "Invalid email address", http.StatusUnauthorized)
		default:
			app.errorLogger.Println("Querying database", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.errorLogger.Println("Comparing password", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !match {
		app.errorLogger.Println("Invalid password")
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// Otherwise, we know the user exists and the password has been checked.
	// Otherwise, if the password is correct, we generate a new token with a 24-hour
	// expiry time and the scope 'authentication'.
	token, err := app.models.Token.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		app.errorLogger.Println("Creating new token", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Encode the token to JSON and send it in the response along with a 201 Created
	// status code.
	response := map[string]interface{}{
		"authentication_token": token}
	responseJson, err := json.Marshal(response)
	if err != nil {
		app.errorLogger.Println("Error marshaling response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}
