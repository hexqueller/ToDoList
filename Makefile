.SILENT:

build:
	CGO_ENABLED=0 GOOS=linux go build -o golang/server golang/main.go

run: build
	docker-compose down && docker-compose up --build

kbuild: build
	minikube image build golang/ -t back:dev
	minikube image build python/ -t front:dev

krun: kbuild
	kubectl delete -f K8S-deployment.yml && kubectl apply -f K8S-deployment.yml