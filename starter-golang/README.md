# Helm / Kubernetes Starter Example:
Local testbed for simple Golang application with a 4 pod ReplicaSet running in a 1 node Kubernetes cluster installed with Helm

## Application
There is a nodejs application running that will respond with the host name it is running on.
There is a LoadBalancer service which exposes the node app server listen port (3000) on port 8080 on the host.

## Installation

### Requirements
Minikube
Docker

### Prepare the environment
We will be pulling a Docker Hub image that is a mere 12Mb in size.
The small size of the image is due to a multistage Docker build that leverages the Alpine Linux build with
a Golang executable that is compiled separately.

## Start the containers
Run the following commands in the specified order:
```
kubectl apply -f kubernetes/namespace.yaml
kubectl apply -f kubernetes/service.yaml
kubectl apply -f kubernetes/deployment.yaml
```

In a separate shell run:

`minikube tunnel` <-- this allows the LoadBalancer service to work in minikube
see: https://minikube.sigs.k8s.io/docs/handbook/accessing/

### Get the _external IP_ address as created by `minikube tunnel`
```
kubectl get svc --namespace=sandkube-ns
```

Note: The use of a namespace here requires all kubectl commands to include the `--namespace=sandkube-ns` argument, or you can install `kubens` to set the active namespace.


## View the Golang application
```
curl <IP Address and port from above>
```
Note it is port 8080 as configured in the LoadBalancer service

Refresh to get responses from the LoadBalancer picked pod. (ctrl-F5, or mouse click refresh + shift)