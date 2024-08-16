.SILENT:

minikubeBuild:
	minikube image build golang/ -t back:dev
	minikube image build python/ -t front:dev

run: minikubeBuild
	helm install project helmChart/

restart:
	helm delete project
	$(MAKE) run