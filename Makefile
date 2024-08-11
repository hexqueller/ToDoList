.SILENT:

build:
	CGO_ENABLED=0 GOOS=linux go build -o golang/server golang/main.go

run: build
	docker-compose down && docker-compose up --build