SOURCES=./
BINS=bin

dep:
	dep ensure

.PHONY: clean
clean:
	rm -rf ${BINS}/dbMigrate

db_migrate:
	go run ${SOURCES}/examples/dbMigrate/main.go

db_new_migration:
	go build -o ${BINS}/newMigration ${SOURCES}/examples/dbNewMigration/main.go
	${BINS}/newMigration Test

.DEFAULT_GOAL := test
test: db_new_migration

