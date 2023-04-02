run:
	go run cmd/server/main.go
	
build:
	go build cmd/server/main.go

docgen:
	swag init -g ./cmd/server/main.go --parseDependency --parseInternal