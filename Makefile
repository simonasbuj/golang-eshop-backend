

run-api:
	@APP_ENV=local go run cmd/rest/main.go

run-server-with-watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec "APP_ENV=local go run cmd/rest/main.go"

run-server-with-watch-wsl:
	nodemon --legacy-watch --watch . --ext go --signal SIGTERM --exec "APP_ENV=local go run cmd/rest/main.go"

run-api-with-air:
	@APP_ENV=local air

test:
	@go test -v ./...


up-docker:
	docker compose up -d

stop-docker:
	docker compose stop

down-docker:
	docker compose down
