package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lucas271/GoSimpleApi/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, "It was not possible to create user, Json error")

	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "claudio",
	})

	if err != nil {
		println(err.Error())
		respondWithError(w, 400, "Couldn't create user")
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))

}
