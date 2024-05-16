swagger:
	swag init -o tmp -d app --parseInternal
	rm -rf tmp

lint:
	golangci-lint run -v

mockery:
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

deploy-check: lint swagger test
