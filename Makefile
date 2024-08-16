.SILENT:

minikubeBuild:
	minikube image build golang/ -t back:dev
	minikube image build python/ -t front:dev

run: minikubeBuild
	helm install project helmChart/

delete:
	helm delete project

restart: delete
	$(MAKE) run