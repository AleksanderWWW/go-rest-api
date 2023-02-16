dev:
	go fmt ./...
	go vet ./...
	go run main.go

test:
	go test -v ./... --cover
