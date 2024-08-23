.SILENT:

minikubeBuild:
	minikube image build golang/ -t hexqueller/k8s-golang-backend:dev
	minikube image build python/ -t hexqueller/k8s-python-frontend:dev

run: minikubeBuild
	helm install project helmChart/

delete:
	helm delete project

restart: delete
	$(MAKE) run

stress:
	kubectl run load-generator-backend --image=busybox -- /bin/sh -c "while true; do wget -q -O- --header='Content-Type: application/json' --post-data='{\"name\": \"testuser\", \"id\": \"0000000001\"}' http://backend:1234/api/create_user; done"
	kubectl run load-generator-frontend --image=busybox -- /bin/sh -c "while true; do wget -q -O- http://frontend/testuser/1440483041; done"

unstress:
	kubectl delete pod load-generator-backend load-generator-frontend

dockerpush:
	docker build golang/ -t hexqueller/k8s-golang-backend
	docker build python/ -t hexqueller/k8s-python-frontend
	docker push hexqueller/k8s-python-frontend
	docker push hexqueller/k8s-golang-backend