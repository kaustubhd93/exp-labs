# Kubernetes Architecture 

## Nodes and Master

- A node is either a physical machine or virtual on which kubernetes is installed. This is where containers are spawned.
- A master is required to orchestrate all of these nodes. This master node exists separately.
- Master and worker nodes both have kubernetes installed.
- Kubernetes can have multi Master-worker clusters to increase high availability.

## Kubernetes components 

- API server:
	- Acts as a front end to kubernetes. Users, devices, cli all talk to the API server.
- Kubelet:
	- It is the agent that runs on each node of the cluster, to ensure that containers are running as expected.
- etcd:
	- A distributed key value store used to store all the information regarding Master, worker nodes and their resources.
- container runtime:
	- Underlying software that is used to run containers. Generally Docker. There are others as well like rkt and CRI-O.
- controller:
	- Monitors if containers go down and trigger actions accordingly. The brain of monitoring inside kubernetes.
- scheduler:
	- Distributes work load.

Official Documentation: https://kubernetes.io/docs/concepts/overview/components/

## Master-worker relationship

- The master has the kube-API server, etcd, controller and scheduler on it's node.
- The worker has container run time and the kubelet agent running.

## kubectl

- The command line utility to control kubernetes components.
- `kubectl cluster-info`

## minikube

- A stand alone single node cluster where all kubernetes services run. In short a sandbox environment.
- `minikube start|stop|status`

## Pods

- It is the smallest deployable unit of computing that can be created and managed in kubernetes.
- A docker container runs inside a POD. Multiple containers can run inside a pod but this should be done only
  if there is a helper container linked to the main application. Multiple containers in a pod are a rare case.
  It is usually when 2 or 3 applications need to be tightly coupled to share storage,namespace and network.
- It is a wrapper around a single container. Kubernetes manages Pods rather than managing the container directly.

Official documentation : https://kubernetes.io/docs/concepts/workloads/pods/

## Pods with yml

- Every YML file will have a default syntax 
```
apiVersion:
kind:
metadata:


spec:
```
- 
