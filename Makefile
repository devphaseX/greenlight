## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: run/api
## run/api: run the cmd/api application
run/api:
	@go run ./cmd/api

.PHONY: db/psql
## db/psql: connect to the database using psql
db/psql:
	psql ${GREENLIGHT_DB_DSN}

.PHONY: db/migrations/new
## db/migrations/new name=$1: create a new database migration
db/migrations/new:
	@echo 'Creating migration files for ${name}'
	migrate create -ext=.sql -seq -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up
