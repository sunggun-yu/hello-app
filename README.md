# release-please-kustomization-bump-demo

This repository demonstrates an automated workflow for updating Docker image tags in k8s manifests. The workflow leverages the following tools:

- **GitHub Actions**
- [**Release Please**](https://github.com/googleapis/release-please): `release-please` creates a new version, and GitHub release with change log.
- [**GoReleaser**](https://goreleaser.com/): A release automation tool for Go projects, used here for building the application and its Docker image.
- [**Kustomization**](https://kustomize.io/): build the k8s manifests and update the Docker image tag using the `kustomize edit` feature.

This repository also provides a simple HTTP service to test various scenarios in the k8s. such as Istio, Service Mesh, etc.

## Usage

### Endpoints

The application exposes the following HTTP endpoints:

- `/`: Root endpoint that returns an HTML page with service information including:
  - Service name
  - Version
  - Instance hostname
  - Host
  - Color
  - Current timestamp

- `/ping`: Health check endpoint that returns a JSON response with:
  - "pong" message
  - Service name
  - Version
  - Instance hostname
  - Current timestamp

- `/hello`: Greeting endpoint that returns a text response
  - Returns a personalized greeting if `NAME` environment variable is set
  - Supports an optional `wait` query parameter (e.g., `/hello?wait=5s` or `/hello?wait=5`) to simulate delay
  - Includes instance hostname in the response
  - Example responses:
    - Basic: `Hello, Instance: hostname`
    - With name: `Hello, John!, Instance: hostname`
    - With wait: `Hello, John!, waited 5s, Instance: hostname`

- `/health`: Simple health check endpoint that returns "HEALTHY" when the service is running

### Environment Variables

The application uses the following environment variables:

- `PORT`: Specifies the port on which the server will listen. Defaults to `8080` if not set. The port is prefixed with `0.0.0.0:` to ensure it listens on all network interfaces.

- `SERVICE`: The name of the service. This is used in the response data for the root (`/`) and `/ping` endpoints.

- `VERSION`: The version of the service. This is also included in the response data for the root (`/`) and `/ping` endpoints.

- `NAME`: The name used in the `/hello` endpoint to personalize the greeting message. If not set, the message defaults to "Hello!".

- `COLOR`: Used in the response data for the root (`/`) endpoint to specify a color, which can be used for theming or display purposes.

These environment variables allow you to configure the behavior and responses of the application without modifying the code.
