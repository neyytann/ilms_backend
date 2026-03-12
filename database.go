package config

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {
    connStr := "postgres://postgres:yourpassword@localhost:5432/ilms_db"
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to connect: %v", err)
    }
    return conn, nil
}
