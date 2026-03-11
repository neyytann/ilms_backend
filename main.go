package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Define the struct BEFORE using it
type Intern struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Preferences json.RawMessage `json:"preferences"`
	Tags string `json:"tags"`
}

func main() {
	connStr := "postgres://postgres:it'smenathan@localhost:5432/ilms_database?sslmode=disable"

	// Connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close(context.Background())

	fmt.Println("Database connected successfully!")

	// Sample data
	prefs := map[string]interface{}{
		"theme":         "dark",
		"notifications": true,
	}
	prefsJSON, _ := json.Marshal(prefs)

	intern := Intern{
		Name: "Nathan",
		Email: "nathan@example.com",
		Preferences: prefsJSON,
		Tags: "golang,backend,intern",
	}

	// Insert into database
	_, err = conn.Exec(
		context.Background(),
		"INSERT INTO interns (name, email, preferences, tags) VALUES ($1, $2, $3, $4)",
		intern.Name, intern.Email, intern.Preferences, intern.Tags,
	)

	if err != nil {
		fmt.Println("Insert failed:", err)
		return
	}

	fmt.Println("Intern inserted successfully!")
}