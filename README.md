# Kube Playground

## Kubernetes, ArgoCD, and Helm Example Repository

This repository contains example Kubernetes YAML files, ArgoCD application manifests, and Helm charts for beginners.

## Contents

- [k8s](./k8s/README.md): Example k8s YAML files.
- [ArgoCD](./ArgoCD/README.md): Example ArgoCD application manifests.
- [Helm](./Helm/README.md): Example Helm charts and values files.

## Usage

Clone this repository and explore the directories for examples.
The Helm chart provides a simple example for deploying an NGINX application with configurable replica count, image, and service settings. You can customize it further to fit your specific application requirements.

## Aliases

```BASH
alias k=kubectl
alias kdy="kubectl --dry-run=client -o yaml"
```

## Upcoming

- Observability examples
