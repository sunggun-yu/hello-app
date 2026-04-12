# Hello App

This repository demonstrates an automated workflow for updating Docker image tags in k8s manifests. The workflow leverages the following tools:

- **GitHub Actions**
- [**Release Please**](https://github.com/googleapis/release-please): `release-please` creates a new version, and GitHub release with change log.
- [**GoReleaser**](https://goreleaser.com/): A release automation tool for Go projects, used here for building the application and its Docker image.
- [**Kustomization**](https://kustomize.io/): build the k8s manifests and update the Docker image tag using the `kustomize edit` feature.

This repository also provides a simple HTTP and GRPC service to test various scenarios in the k8s. such as Istio, Service Mesh, etc.

## Usage

### Web Endpoints

The application exposes the following HTTP endpoints on both the primary (`PORT`, default `8080`) and secondary (`PORT_2`, default `3000`) web servers:

- `/`: Root endpoint that returns an HTML page with service information including:
  - Service name and version
  - Instance hostname
  - Host and server port
  - Client IP and remote address
  - Current timestamp
  - Full HTTP request headers (useful for inspecting Istio/proxy headers)

- `/ping`: Health check endpoint that returns a JSON response with:
  - "pong" message
  - Service name
  - Version
  - Instance hostname
  - Current timestamp
  - Client IP and remote address

- `/hello`: Greeting endpoint that returns a text response
  - Returns a personalized greeting if `NAME` environment variable is set
  - Supports an optional `wait` query parameter (e.g., `/hello?wait=5s` or `/hello?wait=5`) to simulate delay
  - Includes instance hostname in the response
  - Example responses:
    - Basic: `Hello, Instance: hostname`
    - With name: `Hello, John!, Instance: hostname`
    - With wait: `Hello, John!, waited 5s, Instance: hostname`

- `/health`: Readiness check endpoint that returns "HEALTHY" when the service is running

- `/livez`: Liveness check endpoint that returns "OK" unconditionally (used by Kubernetes liveness probe)

### Environment Variables

The application uses the following environment variables:

- `PORT`: Port number for the primary web server. Defaults to `8080`.

- `PORT_2`: Port number for the secondary web server. Defaults to `3000`. Must be different from `PORT`. Application will not start if `PORT` and `PORT_2` have same port. Contains identical endpoints to the primary server. Used to simulate multiple servers in a single Kubernetes Pod.

- `GRPC_PORT`: Port number for the gRPC server. Defaults to `9090`.

- `SERVICE`: The name of the service. This is used in the response data for the root (`/`) and `/ping` endpoints.

- `VERSION`: The version of the service. This is also included in the response data for the root (`/`) and `/ping` endpoints.

- `NAME`: The name used in the `/hello` endpoint to personalize the greeting message. If not set, the message defaults to "Hello!".

- `COLOR`: Used as the accent color for the root (`/`) endpoint UI (header bar background). Accepts any CSS color value (e.g., `#ef476f`, `green`). Defaults to `#7cc423`.

### GRPC Endpoints

#### Describe the HelloService

```bash
grpcurl -plaintext localhost:9090 describe api.HelloService
```

```protobuf
service HelloService {
  rpc Ping ( .api.PingRequest ) returns ( .api.PingResponse );
  rpc SayHello ( .api.HelloRequest ) returns ( .api.HelloReply );
}
```

#### Ping

```bash
grpcurl -plaintext localhost:9090 api.HelloService.Ping
```

result:

```json
{
  "message": "pong",
  "instance": "your-host-name",
  "timestamp": "2024-12-23T14:22:38-05:00"
}
```

#### SayHello

```bash
grpcurl -plaintext -d '{"name": "Darth Vader"}' \
  localhost:9090 api.HelloService.SayHello
```

result:

```json
{
  "message": "Hello Darth Vader"
}
```

## Deployment

### Quick Deploy

Deploy the base configuration to a Kubernetes cluster:

```bash
kubectl create ns hello
kubectl apply -f manifests/install.yaml
```

### Deploy with Kustomize

```bash
# Base deployment
kustomize build manifests/base | kubectl apply -f -

# Service variants (different colors/labels for testing)
kustomize build manifests/overlays/service1 | kubectl apply -f -
kustomize build manifests/overlays/service2 | kubectl apply -f -
kustomize build manifests/overlays/service3 | kubectl apply -f -
```

### Local Development with Minikube

Build and test locally without pushing to a registry:

```bash
# Start minikube
minikube start --container-runtime=containerd --cpus 4 --memory 8192mb

# Build the Go binary for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o hello-app .

# Build the image inside minikube
minikube image build -t hello-app:local .

# Create namespace and deploy
kubectl create ns hello
cd manifests/base && kustomize edit set image hello=hello-app:local && cd ../..
kustomize build manifests/base | kubectl apply -f -

# Verify deployment
kubectl -n hello rollout status deployment/hello
kubectl -n hello get pods

# Test via port-forward
kubectl -n hello port-forward svc/hello 8080:8080 3000:3000 9090:9090

# In another terminal
curl http://localhost:8080/ping
curl http://localhost:8080/health
curl http://localhost:8080/livez
curl http://localhost:8080/hello
grpcurl -plaintext localhost:9090 api.HelloService.Ping
```

### Kubernetes Manifest Details

The deployment includes:

- **ServiceAccount**: `hello`
- **Service**: ClusterIP exposing ports 8080 (HTTP), 3000 (HTTP), 9090 (gRPC)
- **Deployment**: Single replica with:
  - Liveness probe on `/livez` (port 8080)
  - Readiness probe on `/ping` (port 8080)
  - Resource limits: 200m CPU, 256Mi memory
  - Resource requests: 20m CPU, 128Mi memory
  - Graceful shutdown support (handles SIGTERM)

### Service Overlays

Three overlay variants are provided for testing multi-service scenarios (e.g., Istio traffic routing):

| Overlay    | Color     | Labels                    |
| ---------- | --------- | ------------------------- |
| `service1` | `#ef476f` | service=Service1, version=V2 |
| `service2` | `#ffd166` | service=Service2, version=V2 |
| `service3` | `#06d6a0` | service=Service3, version=V1 |
