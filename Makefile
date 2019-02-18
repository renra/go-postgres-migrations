SOURCES=./
BINS=bin

dep:
	dep ensure

.PHONY: clean
clean:
	rm -rf ${BINS}/migrate
	rm -rf ${BINS}/newMigration

db_migrate:
	go build -o ${BINS}/migrate ${SOURCES}/examples/dbMigrate/main.go
	${BINS}/migrate
	${BINS}/migrate down

db_new_migration:
	go build -o ${BINS}/newMigration ${SOURCES}/examples/dbNewMigration/main.go
	${BINS}/newMigration Test

.DEFAULT_GOAL := test
test: db_new_migration db_migrate

