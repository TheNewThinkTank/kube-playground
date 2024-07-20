# Helm

[Helm](https://helm.sh/)

```BASH
helm list -A
helm install
helm uninstall
helm status
helm get -h
helm -h
helm version
helm env  # see environment variables
helm -debug  # for verbose output

helm search hub [chart-name]
helm repo add [chart-name] [chart-url]
helm search repo [chart-name]
helm repo list

helm install [release-name] [chart name]

helm list
helm uninstall [release-name]
helm pull --untar [chart-name]
ls [chart-name]
```
