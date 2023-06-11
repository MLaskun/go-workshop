build:
	@go build -o bin/go-workshop
run: build
	@./bin/go-workshop
test:
	@go test -v ./...