.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o golang/server golang/main.go

minikubeBuild: build
	minikube image build golang/ -t back:dev
	minikube image build python/ -t front:dev

run: minikubeBuild
	kubectl apply -f .

restart: minikubeBuild
	kubectl delete -f . && kubectl apply -f .