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
	Name      string         `json:"name"`
	Email     sql.NullString `json:"email"`
	Password  string         `json:"password"` // Change to `json:"-"`
	ApiKey    string         `json:"api_key"`  // Change to `json:"-"`
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
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.HashedPassword,
		ApiKey:    user.ApiKey,
	}, nil
}
