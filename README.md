# Postgres Migrations

A tool that simplifies working with postgres migrations building on [sql-migrate](https://github.com/rubenv/sql-migrate)

## Usage

### Generating migrations

```go
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
```

Then you only need to implement the `Up` and `Down` sections of your migration and then ...


### Migrating

```go
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
```

If you build this as a program `migrate` you can call `migrate` to run all pending migrations and `migrate down` to roll back 1 mig at a time.
