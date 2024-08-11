.SILENT:

build:
	CGO_ENABLED=0 GOOS=linux go build -o server main.go

docker: build
	docker build . -t back

run: docker
	docker run -p 1234:1234 back