# Helm / Kubernetes Starter Example:
Local testbed for simple NodeJS application with a 2 pod ReplicaSet running in a 1 node Kubernetes cluster installed with Helm

## Application
There is a nodejs application running that will respond with the host name it is running on.
There is a LoadBalancer service which exposes the node app server listen port (3000) on port 8080 on the host.

## Installation

### Requirements
Minikube
Docker
Helm

### Prepare the environment
We will be building local Docker images and pulling them in directly.
To accomplish this with Minikube you must set your current shell to use the Minikube docker context.
Run:

```
eval $(minikube docker-env)
docker build -t sandkube .
```

## Start the Kubernetes cluster with Helm
```
helm install sandkube temp/sandcube
```

In a separate shell run:

`minikube tunnel` <-- this allows the LoadBalancer service to work in minikube
see: https://minikube.sigs.k8s.io/docs/handbook/accessing/

### Get the _external IP_ address as created by `minikube tunnel`
```
kubectl get svc --namespace=sandkube-ns
```

## View the NodeJS application
```
curl <IP Address from above>
```
Note it is port 8080 as configured in the LoadBalancer service

Refresh to get responses from the LoadBalancer picked pod. (ctrl-F5, or mouse click refresh + shift)