package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SumDeusVitae/cli-assistant/internal/database"
	"github.com/SumDeusVitae/cli-assistant/internal/gpt"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUserRequest(w http.ResponseWriter, r *http.Request, user database.User) {
	params := struct {
		Model   string `json:"model"`
		Request string `json:"request"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	id := uuid.New().String()
	err = cfg.DB.CreateComun(r.Context(), database.CreateComunParams{
		ID:        id,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Model:     params.Model,
		Question:  params.Request,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create note")
		return
	}

	// here we are making request to model and returning reply
	replyMessage, err := gpt.RequestGPT(params.Request, cfg.GptApiKey)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusExpectationFailed, "Couldn't get reply from ai")
		return
	}
	// log.Println(replyMessage)
	// updating db with reply
	reply := sql.NullString{String: replyMessage, Valid: true}

	err = cfg.DB.UpdateReply(r.Context(), database.UpdateReplyParams{
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Reply:     reply,
		ID:        id,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't update reply in db")
		return
	}
	// and sending reply back

	communication, err := cfg.DB.GetComunsById(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get communication")
		return
	}

	comunResp, err := databaseComuntoComun(communication)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't convert communication")
		return
	}

	respondWithJSON(w, http.StatusCreated, comunResp)

}
