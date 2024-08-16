.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o dev-server main.go

run: build
	./dev-server