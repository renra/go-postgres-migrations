package main

import (
  "os"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
  "app/dbMigrate"
)

type Logger struct {
}

func (l *Logger) Log(msg string) {
  fmt.Println(msg)
}

func getDirection() string {
  if len(os.Args) > 1 {
    return os.Args[1]
  } else {
    return ""
  }
}

func main() {
  logger := &Logger{}

  connString := fmt.Sprintf(
    "user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
    os.Getenv("POSTGRES_USER"),
    os.Getenv("POSTGRES_PASSWORD"),
    os.Getenv("POSTGRES_DATABASE"),
    os.Getenv("POSTGRES_HOST"),
    os.Getenv("POSTGRES_PORT"),
  )

  conn, err := sql.Open(dbMigrate.DBType, connString)

  if err != nil {
    panic(err)
  }

  migrationsDir := "migrations"

  dbMigrate.Run(conn, getDirection(), migrationsDir, logger)
}
