package database

import (
	"database/sql"
	"finalProject4/entity"
	"finalProject4/pkg/helpers"
	"log"
)

func SeedAdmin(db *sql.DB) {
	password, _ := helpers.GenerateHashedPassword([]byte("admin123"))
	insertAdminQuery := `
		INSERT INTO users (full_name, email, password, role) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING;
	`
	var admin = entity.User{
		FullName: "admin",
		Email:    "admin@mail.com",
		Password: password,
		Role:     "admin",
	}
	err := db.QueryRow(insertAdminQuery, admin.FullName, admin.Email, admin.Password, admin.Role).Err()
	if err != nil {
		log.Fatal(err)
	}
}
