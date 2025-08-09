

run-api:
	@APP_ENV=local go run cmd/rest/main.go

run-server-with-watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec "APP_ENV=local go run cmd/rest/main.go"

run-server-with-watch-wsl:
	nodemon --legacy-watch --watch . --ext go --signal SIGTERM --exec "APP_ENV=local go run cmd/rest/main.go"

test:
	@go test -v ./...