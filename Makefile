build:
	docker build -t k3d-myreg:5000/app1 .
	docker push k3d-myreg:5000/app1
	helm upgrade --install dapractors dapractors

registry:
	docker pull ghcr.io/dapr/dapr:1.14.4
	docker tag ghcr.io/dapr/dapr:1.14.4 k3d-myreg:5000/dapr/dapr:1.14.4
	docker push k3d-myreg:5000/dapr/dapr:1.14.4


	docker pull ghcr.io/dapr/operator:1.14.4
	docker tag ghcr.io/dapr/operator:1.14.4 k3d-myreg:5000/dapr/operator:1.14.4
	docker push k3d-myreg:5000/dapr/operator:1.14.4

	docker pull ghcr.io/dapr/placement:1.14.4
	docker tag ghcr.io/dapr/placement:1.14.4 k3d-myreg:5000/dapr/placement:1.14.4
	docker push k3d-myreg:5000/dapr/placement:1.14.4

	docker pull ghcr.io/dapr/sentry:1.14.4
	docker tag ghcr.io/dapr/sentry:1.14.4 k3d-myreg:5000/dapr/sentry:1.14.4
	docker push k3d-myreg:5000/dapr/sentry:1.14.4

	docker pull ghcr.io/dapr/injector:1.14.4
	docker tag ghcr.io/dapr/injector:1.14.4 k3d-myreg:5000/dapr/injector:1.14.4
	docker push k3d-myreg:5000/dapr/injector:1.14.4

	docker pull ghcr.io/dapr/scheduler:1.14.4
	docker tag ghcr.io/dapr/scheduler:1.14.4 k3d-myreg:5000/dapr/scheduler:1.14.4
	docker push k3d-myreg:5000/dapr/scheduler:1.14.4

	docker pull ghcr.io/dapr/daprd:1.14.4
	docker tag ghcr.io/dapr/daprd:1.14.4 k3d-myreg:5000/dapr/daprd:1.14.4
	docker push k3d-myreg:5000/dapr/daprd:1.14.4

cluster-rebuild:
	k3d cluster delete
	k3d cluster create --registry-use k3d-myreg:5000 --api-port 6550 -p "8081:80@loadbalancer" --agents 2
registry-create:
	k3d registry create myreg -p 5000
	k3d cluster create --registry-use k3d-myreg:5000 --api-port 6550 -p "8081:80@loadbalancer"
