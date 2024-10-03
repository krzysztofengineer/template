export GOOSE_DRIVER=sqlite3
export GOOSE_DBSTRING=./db/db.sqlite
export GOOSE_MIGRATION_DIR=./db/migrations

dev: migrate 
	air

build: templ tailwind sqlc
	go build -o ./tmp/main .

test: templ tailwind sqlc
	go test ./...

templ:
	templ generate

tailwind:
	tailwindcss -i css/main.css -o dist/main.css

create_migration:
	goose -s create $(name) sql

migrate:
	goose up

sqlc:
	sqlc generate