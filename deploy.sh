docker build -t carlo/downfile:v1 .
kubectl apply -f downFile.yaml -n prod
kubectl apply -f downFile-gateway.yaml -n prod