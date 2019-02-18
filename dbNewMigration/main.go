package dbNewMigration

import (
  "fmt"
  "time"
  "io/ioutil"
  i "app/interfaces"
)

func Run(name string, migrationsDir string, logger i.Logger) {
  // no idea why this number, but it outputs the current timestamp
  timeStamp := time.Now().Format("20060102150405")

  path := fmt.Sprintf("%s/%s_%s.sql", migrationsDir, timeStamp, name)

  contents := []byte("-- +migrate Up\nCREATE TABLE somethings (id SERIAL PRIMARY KEY);\n\n-- +migrate Down\nDROP TABLE somethings;")

  err := ioutil.WriteFile(path, contents, 0644)

  if err != nil {
    panic(err)
  }

  logger.Log(fmt.Sprintf("Check your new migration here: %s", path))
}

