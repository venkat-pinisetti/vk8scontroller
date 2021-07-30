# Custom Kubernetes Controller
This sample custom kubernetes controller.

## Pre-requisites:
Go installation & configuration
Venilla Kubernetes cluster

## Steps:
1. Git clone the repo
$ git clone https://github.com/venkat-pinisetti/vk8scontroller.git

2. Change directory to cloned path
$ cd vk8scontroller.git

## How to run
1. make build
2. export kube config
3. run vcontroller

Example:

root@master11:~/go/src/k8s.io/vk8scontroller# make build
go build -o vcontroller ./cmd/main.go
root@master11:~/go/src/k8s.io/vk8scontroller# ll vcontroller 
-rwxr-xr-x 1 root root 37893833 Jul 30 03:13 vcontroller*
root@master11:~/go/src/k8s.io/vk8scontroller# export KUBECONFIG=/root/.kube/config
root@master11:~/go/src/k8s.io/vk8scontroller# ./vcontroller 
I0730 03:45:48.831060   14546 main.go:16] Setting Kube client
I0730 03:45:48.834896   14546 main.go:31] Creating controller
I0730 03:45:48.835088   14546 controller.go:36] Setting up event handlers
I0730 03:45:48.835184   14546 controller.go:45] Starting Zbm controller
I0730 03:45:48.835200   14546 controller.go:46] Waiting for informer caches to sync
I0730 03:45:48.935576   14546 controller.go:50] Starting workers
I0730 03:45:48.935616   14546 controller.go:52] Started workers

## References
https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#:~:text=Custom%20controllers&text=When%20you%20combine%20a%20custom,sync%20with%20the%20desired%20state.

https://github.com/kubernetes/sample-controller

