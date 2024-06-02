swagger:
	swag init -o app/docs -d app --parseInternal

check-swagger:
	swag init -o tmp -d app --parseInternal
	rm -rf tmp

lint:
	brew upgrade golangci-lint
	golangci-lint run -v

mockery:
	rm -rf mocks
	mockery

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
