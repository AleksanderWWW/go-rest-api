dev:
	go fmt ./...
	go vet ./...
	go run httpd/main.go
