swagger:
	swag init -o app/docs -d app --parseInternal

check-swagger:
	swag init -o tmp -d app --parseInternal
	rm -rf tmp

easyjson:
	 easyjson app/pkg/log/jsonHandler.go

update-linter:
	brew upgrade golangci-lint

lint: easyjson
	golangci-lint run -v

mockery: easyjson
	find . -type f -name 'mock_*' -exec rm {} +
	find . -type f -name "mockWrappers.go" -execdir mv -n -- {} mockWrappers.txt \; # Костылина, чтобы не падала генерация моков
	mockery
	find . -type f -name "mockWrappers.txt" -execdir mv -n -- {} mockWrappers.go \;

test-coverage-number: mockery
	go test -v -coverprofile=profile.cov ./app...
	go tool cover -func profile.cov
	rm profile.cov

test-coverage-html: mockery
	go test -v -coverprofile=profile.cov ./app...
	go tool cover -html profile.cov
	rm profile.cov

test: mockery
	go test ./...

deploy-check: check-swagger test lint

migration-create-file:
	goose -dir migrations/pgsql create init sql

migration-up:
	goose -dir migrations/pgsql up

migration-down:
	goose -dir migrations/pgsql down
