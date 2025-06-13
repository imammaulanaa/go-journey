# Journey 19: Deployment ke Docker & Kubernetes
> **Cerita**  
Gopher menaiki armada Docker multi-stage, membangun image ringan, lalu membentangkan layar Kubernetes. Ia mengaplikasikan manifest Deployment dan Service, membangun Ingress gate, dan menyebar armadanya ke pulau cloud.


## Tujuan Pembelajaran
- Dockerfile multi-stage
- docker-compose.yml
- Kubernetes manifests

## Daftar File
- `Dockerfile`
- `docker-compose.yml`
- `k8s/deployment.yaml`
- `k8s/service.yaml`
- `k8s/ingress.yaml`

## Cara Menjalankan
```bash
cd journey-19-deployment-docker-kubernetes
docker build -t go-quest-app . && docker run -p 8080:8080 go-quest-app
```

## Catatan
Untuk K8s: kubectl apply -f k8s/

## Referensi
- https://kubernetes.io
- https://docs.docker.com/develop/develop-images/multistage-build