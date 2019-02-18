package main

import (
  "os"
  "fmt"
  "app/dbNewMigration"
)

type Logger struct {
}

func (l *Logger) Log(msg string) {
  fmt.Println(msg)
}

func main() {
  if len(os.Args) < 2 {
    panic("You had better give me a file name for your migration")
  }

  migrationsDir := "migrations"
  migName := os.Args[1]
  logger := &Logger{}

  dbNewMigration.Run(migName, migrationsDir, logger)
}


