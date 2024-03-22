# release-please-kustomization-bump-demo

This repository demonstrates an automated workflow for updating Docker image tags in k8s manifests. The workflow leverages the following tools:

- **GitHub Actions**
- [**Release Please**](https://github.com/googleapis/release-please): `release-please` creates a new version, and GitHub release with change log.
- [**GoReleaser**](https://goreleaser.com/): A release automation tool for Go projects, used here for building the application and its Docker image.
- [**Kustomization**](https://kustomize.io/): build the k8s manifests and update the Docker image tag using the `kustomize edit` feature.
