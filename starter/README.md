# Kubernetes Starter Example:
Local testbed for simple NodeJS application with a 2 pod ReplicaSet running in a 1 node Kubernetes cluster.

## Application
There is a nodejs application running that will respond with the host name it is running on.
There is a LoadBalancer service which exposes the node app server listen port (3000) on port 8080 on the host.

## Installation

### Requirements
Minikube
Docker

### Prepare the environment
We will be building local Docker images and pulling them in directly.
To accomplish this with Minikube you must set your current shell to use the Minikube docker context.
Run:

```
eval $(minikube docker-env)
docker build -t sandkube .
```

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

## Excersizes

#### View the NodeJS application
```
curl <IP Address from above>
```
Note it is port 8080 as configured in the LoadBalancer service

Refresh to get responses from the LoadBalancer picked pod. (ctrl-F5, or mouse click refresh + shift)

#### Scale up the number of pods in the deployment

```
kubectl scale --replicas=5 deployment/sandkube-app-deployment --namespace=sandkube-ns
```


## Shut down the container (by namespace!)

```
kubectl delete all --all --namespace=sandkube-ns
```

Ctrl-C in the second terminal to stop the `minikube tunnel` process
