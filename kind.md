# Kind

Firstly, ensure that the Docker daemon is running.

```BASH
# add go dir to PATH
export PATH=$PATH:$(go env GOPATH)/bin

go install sigs.k8s.io/kind@v0.23.0
kind create cluster

k cluster-info --context kind-kind

# inspect Kind cluster with kubectl and/or k9s
k get node
k9s

kind delete cluster
```

## argoCD

```BASH
k create ns argocd

k apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml\n
k get po -n argocd
k port-forward svc/argocd-server -n argocd 8080:443\n

# initial argocd username: admin
# getting initial argocd password the hard way
k -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

# .. or the easy way
brew install argocd
argocd admin initial-password -n argocd
```

## harbor

```BASH
k create ns harbor
helm repo add harbor https://helm.goharbor.io
helm repo update
helm install harbor harbor/harbor --namespace harbor
k get po -n harbor
k port-forward svc/harbor-core -n harbor 8081:80
kubectl get svc -n harbor
```

**user**: admin<br>
**default password**: Harbor12345
