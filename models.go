package main

import (
	"time"

	"github.com/lucas271/GoSimpleApi/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		UpdatedAt: dbUser.UpdatedAt,
		CreatedAt: dbUser.CreatedAt,
	}
}
