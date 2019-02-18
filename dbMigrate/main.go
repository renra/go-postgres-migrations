package main

import (
  "os"
  "fmt"
  "app/app/db"
  "github.com/rubenv/sql-migrate"
  "app/app/config"
)

func main() {
  config := config.Load()

  conn := db.Connect(config)
  defer conn.Close()

  migrations := &migrate.FileMigrationSource{
    Dir: config.GetString("migrations_dir"),
  }

  if len(os.Args) > 1 && os.Args[1] == "down" {
    n, err := migrate.ExecMax(conn, "postgres", migrations, migrate.Down, 1)

    if err != nil {
      panic(err)
    }

    if n == 1 {
      fmt.Println("Rolled back successfully")
    } else {
      fmt.Println("Nothing to roll back")
    }

  } else {
    n, err := migrate.Exec(conn, "postgres", migrations, migrate.Up)

    if err != nil {
      panic(err)
    }

    fmt.Printf("Applied %d migrations!\n", n)
  }
}
