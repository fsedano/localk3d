helm repo add dapr https://dapr.github.io/helm-charts/
helm repo update
helm upgrade --install dapr dapr/dapr \
--version=1.14 \
--namespace dapr-system \
--create-namespace \
--wait


k3d registry create myreg -p 5000
k3d cluster create --registry-use k3d-myreg:5000 --api-port 6550 -p "8081:80@loadbalancer"
docker tag nginx:latest k3d-myreg:5000/mynginx:v0.1
docker push k3d-myreg:5000/mynginx:v0.1
kubectl run mynginx --image k3d-myreg:5000/mynginx:v0.1

--api-port 6550 -p "8081:80@loadbalancer"

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
