.SILENT:

build:
	docker build . -t web

run: build
	docker run -p 8080:5000 web