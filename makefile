tests:
	ENVIRONMENT=test go test -count 1 -timeout 7m -p 1 -covermode count -coverprofile code-coverage.out -v ./...

coverage: view
	tests

view:
	go tool cover -html=code-coverage.out


