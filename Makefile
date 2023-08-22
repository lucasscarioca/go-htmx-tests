include .env
export

assets:
	@tailwindcss -i ./assets/tailwind.css -o ./assets/static/styles.css

build: assets
	@go build ./cmd/server/... -o tmp/main

start: build
	@./tmp/main

run: assets
	@go run ./cmd/server/...

test:
	@go test -v ./...

install:
	@npm i -g tailwindcss
	@go mod tidy

# create_migration:
# 	migrate create -ext sql -dir configs/migrations/ -seq $(name)

# migration_up:
# 	migrate -path configs/migrations/ -database ${DB_URL} -verbose up

# migration_down:
# 	migrate -path configs/migrations/ -database ${DB_URL} -verbose down

# migration_fix:
# 	migrate -path configs/migrations/ -database ${DB_URL} force $(version)
