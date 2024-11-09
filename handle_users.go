package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SumDeusVitae/cli-assistant/internal/auth"
	"github.com/SumDeusVitae/cli-assistant/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type Email struct {
		String string `json:"email"`
		Valid  bool   `json:"valid"` // Valid is true if String is not NULL
	}
	type parameters struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Email    Email  `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	hashed_password, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("Couldn't hash password: %v", err)
	}

	apiKey, err := generateRandomSHA256Hash()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't gen apikey")
		return
	}

	var email sql.NullString
	if params.Email.Valid && params.Email.String != "" {
		email = sql.NullString{String: params.Email.String, Valid: true}
	} else {
		// Set the email to NULL if invalid or empty
		email = sql.NullString{Valid: false}
	}

	err = cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New().String(),
		CreatedAt:      time.Now().UTC().Format(time.RFC3339),
		UpdatedAt:      time.Now().UTC().Format(time.RFC3339),
		Login:          params.Login,
		Email:          email,
		HashedPassword: hashed_password,
		ApiKey:         apiKey,
	})
	if err != nil {
		// Check if the error is a UNIQUE constraint violation (specific to SQLite)
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println("Error: UNIQUE constraint failed")
			respondWithError(w, http.StatusConflict, "User already exists") // 409 Conflict
		} else {
			log.Println("Error inserting user:", err)
			respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		}
		return
	}
	user, err := cfg.DB.GetUser(r.Context(), apiKey)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}
	userResp, err := databaseUserToUser(user)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't convert user")
		return
	}
	respondWithJSON(w, http.StatusCreated, userResp)
}

func (cfg *apiConfig) handlerUsersLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	user, err := cfg.DB.GetUserByLogin(r.Context(), params.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			respondWithError(w, http.StatusUnauthorized, "Incorrect login or password")
			return
		}
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't receive data from db")
		return
	}
	err = auth.CheckPasswordHash(user.HashedPassword, params.Password)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Incorrect login or password")
	}
	userResp, err := databaseUserToUser(user)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't convert user")
		return
	}
	respondWithJSON(w, http.StatusCreated, userResp)

}

func generateRandomSHA256Hash() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(randomBytes)
	hashString := hex.EncodeToString(hash[:])
	return hashString, nil

}

func (cfg *apiConfig) resetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, "Forbiden")
	}
	err := cfg.DB.DeleteUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete users")
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Everything reset")); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (cfg *apiConfig) handlerUserGet(w http.ResponseWriter, r *http.Request, user database.User) {
	userResp, err := databaseUserToUser(user)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't convers user")
		return
	}
	respondWithJSON(w, http.StatusOK, userResp)
}
