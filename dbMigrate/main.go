package dbMigrate

import (
  "os"
  "fmt"
  "database/sql"
  "github.com/rubenv/sql-migrate"
)

const (
  Down = "down"
  DBType = "postgres"
)

func GetDirection() string {
  if len(os.Args) > 1 {
    return os.Args[1]
  } else {
    return ""
  }
}

type Logger interface {
  Log(string)
}

func Run(conn *sql.DB, direction string, migrationsDir string, logger Logger) {
  migrations := &migrate.FileMigrationSource{
    Dir: migrationsDir,
  }

  if direction == Down {
    n, err := migrate.ExecMax(conn, DBType, migrations, migrate.Down, 1)

    if err != nil {
      panic(err)
    }

    if n == 1 {
      logger.Log("Rolled back successfully")
    } else {
      logger.Log("Nothing to roll back")
    }

  } else {
    n, err := migrate.Exec(conn, DBType, migrations, migrate.Up)

    if err != nil {
      panic(err)
    }

    logger.Log(fmt.Sprintf("Applied %d migrations!\n", n))
  }
}
