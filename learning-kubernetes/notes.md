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

> Official Documentation : https://kubernetes.io/docs/concepts/workloads/pods/

### Pods with yml

- Every YML file will have a default syntax 
```
apiVersion: v1
kind: Pod
metadata:


spec:
    containers:
```

## Replication Controllers and Replica sets.

### Replication controllers

- It is like a process supervisor, instead of supervising individual processes on a single node, it supervises multiple pods across multiple nodes.
- If there are too many pods, the ReplicationController terminates the extra pods. If there are too few, the ReplicationController starts more pods.
- Few and extra in this case means that there is a proper value set for how many replicas should be running. If it is set to 3 the ReplicationController
  will make sure that only 3 are running and if one or more fails it will restart the failed ones and bring the count to 3 again.
- However Kubernetes recommends to use ReplicaSets over ReplicationControllers.

> Official Documentation : https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/

### Replica Sets 

- It functions similar to Replication Controllers except has a different `apiVersion` which is `apps/v1`.
- There is another flag called `selector`. A ReplicaSet identifies new Pods to acquire by using its selector. If there is a Pod that has no OwnerReference or the OwnerReference is not a Controller and it matches a ReplicaSet's selector, it will be immediately acquired by said ReplicaSet.
- So if there was a pod launched without a replicaset, the selector using `matchLabels` will match the appropriate labels and the replicaset will acquire this pod under it's own control. For example if you launched a replicaset with 3 nginx replicas, and one pod with nginx with one label matching is already running. Replicaset will identify this and launch only 2 more pods keeping the count to desired which is 3.

> Official Documentation : https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/
> NOTE: Replication sets and Replication controllers come under the category of Controllers.

## Deployment

- You can describe a desired state in a deployment and the deployment controller will change the actual state to a desired state.
- This is generally done by defining deployments to create Replicasets.
- Any new changes in the resource defined in the Yaml, deployment controller will create a new set of replica first and then terminate the old one.
- We can even undo the most recent deployment by using the `--record` option while deploying the first time itself. This means, you deploy with `--record` option which starts storing it's rollout history. This is what enables us to rollback to a previous version easily.

> Official Documentation : https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

## Kubernetes Network Model

- Every pod gets it's own IP. This removes the need to create links between pods. We almost never have to deal with mapping container ports to host ports. This creates a model where pods can be treated as VMs from the perspective of port allocation, naming, service discovery, load balancing, application configuration, and migration.
- Kubernetes imposes the following fundamental requirements on any networking implementation (barring any intentional network segmentation policies):
	- pods on a node can communicate with all pods on all nodes without NAT
	- agents on a node (e.g. system daemons, kubelet) can communicate with all pods on that node
	- For those platforms that support Pods running in the host network (e.g. Linux): pods in the host network of a node can communicate with all pods on all nodes without NAT

> Official Documentation : https://kubernetes.io/docs/concepts/cluster-administration/networking/

## Services

- It is an abstract way to expose an application running on a set of pods as a network service.
- The set of pods targeted by a service is usually determined by a selector. 
- Just add the labels of `metadata` to the `selector:` section and your pods are mapped to the service.

### NodePort

- Also services are actually used to connect to pods from the outside. We can access the pod via the nodeport which is exposed on the Cluster. The service balances the load among the pods in a random manner. It can serve as a load balancer.

> Official Documentation : https://kubernetes.io/docs/concepts/services-networking/service/

### ClusterIP

- ClusterIP acts as a single entry point for pods that you don't want to expose to the world and want to use inside the cluster only. 
- A clusterIP will balance the load among the mapped pods in a random manner.
- This kind is usually used while working with databases or APIs that you dont want to expose to the world and use them internally.

> NOTE: While writing a service, inside metadata the name field should match the hostname you want to connect it to. For example if you are connecting to a database with hostname `redis` then the `name:` field inside the `metadata:` section should be `redis`. If this is not the same it causes dns resolution failure and your app wont be able to connect to the database. 

