package main

import (
	"database/sql"
	"time"

	"github.com/SumDeusVitae/cli-assistant/internal/database"
)

type User struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Login     string         `json:"login"`
	Email     sql.NullString `json:"email"`
	Password  string         `json:"-"` // `json:"password"`
	ApiKey    string         `json:"api_key"`
}

func databaseUserToUser(user database.User) (User, error) {
	createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	updatedAt, err := time.Parse(time.RFC3339, user.UpdatedAt)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:        user.ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Login:     user.Login,
		Email:     user.Email,
		Password:  user.HashedPassword,
		ApiKey:    user.ApiKey,
	}, nil
}

type Communication struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Model     string         `json:"model"`
	Question  string         `json:"question"`
	Reply     sql.NullString `json:"reply"`
	UserID    string         `json:"user_id"`
}

func databaseComuntoComun(post database.Communication) (Communication, error) {
	createdAt, err := time.Parse(time.RFC3339, post.CreatedAt)
	if err != nil {
		return Communication{}, err
	}
	updatedAt, err := time.Parse(time.RFC3339, post.UpdatedAt)
	if err != nil {
		return Communication{}, err
	}

	return Communication{
		ID:        post.ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Model:     post.Model,
		Question:  post.Question,
		Reply:     post.Reply,
		UserID:    post.UserID,
	}, nil
}
