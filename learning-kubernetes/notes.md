# Kubernetes Architecture 

## Nodes and Master

- A node is either a physical machine or virtual on which kubernetes is installed. This is where containers are spawned.
- A master is required to orchestrate all of these nodes. This master node exists separately.
- Master and worker nodes both have kubernetes installed.
- Kubernetes can have multi Master-worker clusters to increase high availability.

## Kubernetes components 

- API server:
	- Acts as a front end to kubernetes. Users, devices, cli all talk to the API server.
- etcd:
	- A distributed key value store used to store all the information regarding Master, worker nodes and their resources.
- container runtime:
	- Underlying software that is used to run containers. Generally Docker. There are others as well like containerd, rkt and CRI-O.
- controller:
	- Monitors if containers go down and trigger actions accordingly. The brain of monitoring inside kubernetes.
- scheduler:
	- Distributes work load.
- kube-proxy:
	- runs on all worker nodes, enables communication between services within the cluster by maintaining network rules on worker nodes. These network rules allow network communication to your pods from network sessions inside or outside of your cluster.
- Kubelet:
	- It is the agent that runs on each node of the cluster, to ensure that containers are running as expected.

Official Documentation: https://kubernetes.io/docs/concepts/overview/components/

## Master-worker relationship

- The master has the kube-API server, etcd, controller and scheduler on it's node.
- The worker has container run time, kube-proxy and the kubelet agent running.

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

> NOTE: While writing a service, inside metadata the name field should match the hostname you want to connect it to. For example if you want to connect to a database with hostname `redis` then the `name:` field inside the `metadata:` section should be `redis`. If this is not the same it causes dns resolution failure and your app wont be able to connect to the database. 

## etcd

- etcd stores all information regarding all componenets of kubernetes and the ones you deploy. 
- If k8s cluster was created manually, it is deployed as a systemd service. If the cluster was deployed using kubeadm then it is deployed as a pod in the kube-system namespace.
- Multiple control planes ( multiple masters) will have multiple etcd running in HA.

## scheduler

- It decides which pod should be scheduled on which node by looking at the resource requirements of the pod if defined. 
- It will first short list the node which will be the best fit for the pod and then the kubelet will start the pod on the decided node. 
- If there is no scheduler running in k8s, we can manually schedule the pod by specifying the `nodeName` parameter. The Apiserver then creates a binding object for the same

## Namespace

- It helps in isolating components like deployments, pods, services, etc. We can use the same HA kubernetes cluster and isolate dev and prod envs using namespaces. 
- Namespace-based scoping is applicable only for namespaced objects (e.g. Deployments, Services, etc) and not for cluster-wide objects (e.g. StorageClass, Nodes, PersistentVolumes, etc).
- In order to access any service in another namespace use the domain `<service-name>.<namespace>.svc.cluster.local`. 

## Taints and Tolerations

- Let's say if we want to restrict our nodes to run only a specific kind of pod which requires either heavy or light resources, in this case we can `taint` the node with a key-value pair describing what kind of pods will run here. For eg: `app=frontend`. The `taint` command also requires the effect which denotes what the node will do with an incoming pod request. 
- `kubect taint node node-name <key1>=<value1>:<effect>`
- `<effect>` can have 3 values : NoSchedule, NoExecute, PreferNoSchedule
- Tolerations are applied to pods in the below format. Just because a toleration is applied to the pod that doesnt mean it will be scheduled on the tainted node only. It will be scheduled wherever the scheduler wants to schedule according to the scheduling algorithm.
```
tolerations:
- key: "key1"
  operator: "Equal"
  value: "value1"
  effect: "NoSchedule"
```
- Untainting a node : `kubectl taint nodes node-name key=value:effect-`. Note the `-` at the end, it means to untaint. 

## Node Affinity

- Imagine your pod requires a lot of resources to run and the nodes available are a combination of large,medium and small w.r.t resources. Using taints and tolerations will not garuntee the placement of the pod. In this case we can use node affinity to make sure these pods are properly placed on the appropriate node. 
- Node Affinity has 2 types:
	- `nodeSelector`
	- `nodeAffinity`
- `nodeSelector` provides a very simple way to restrict pods to nodes with particular labels. `nodeAffinity` is conceptually similar to `nodeSelector` but `nodeAffinity` allows users to use excpressions. For eg: If you dont want the pod to be scheduled on a small node and you are okay with it being scheduled on either a large or a medium node we can use nodeaffinity expressions like this.
```
		nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: size
                operator: NotIn
                values:
                - small
```

## DaemonSet

- A daemonset ensures that all nodes run one copy of the pod. Use case: kube-proxy, fluentd log shipper, etc.
- The definition of daemonset is similar to ReplicaSet without the `replicas` parameter.

## Static Pods

- Do you ever wonder who his monitoring the kube api-server, controllermanager, scheduler and etcd ? It is actually kubelet that makes sure all these components are always up. 
- These are static pods. 
- The kubelet service can deploy pods directly without the api-server. Keep the manifests in a directory(usually /etc/kubernetes/manifests) which the kubelet is configured to read from and kubelet will deploy the static pod on that particular node.
- kubelet also makes a mirror object which appears in the `kubectl get pods` result. However, you cannot control the pod from kubectl.

## Multiple schedulers

- Sometimes the default scheduler is not enough because your k8s workload may require a different scheduling method. But Why ? > https://www.cncf.io/blog/2020/08/10/why-the-kubernetes-scheduler-is-not-enough-for-your-ai-workloads/
- We can use a custom scheduler and configure the pod to use this custom scheduler. Add below config to your pod definition file under spec.
```
    shedulerName:
```

## ConfigMaps

> For plugging in the entire configMap as env variables use below yaml config:
```
envFrom:
  - configMapRef:
      name: webapp-config-map
```

## Multiple containers

- A pod can encapsulate an application composed of multiple co-located containers that are tightly coupled and need to share resources. 
- Pods natively provide 2 kinds of shared resources for their containers: 
    - network
    - storage
- These co-located containers form a single cohesive unit of service—for example, one container serving data stored in a shared volume to the public, while a separate sidecar container refreshes or updates those files. The Pod wraps these containers, storage resources, and an ephemeral network identity together as a single unit.
- Refer https://kubernetes.io/docs/concepts/workloads/pods/#how-pods-manage-multiple-containers
- Pods also have initContainers. As the name suggests these are run before the main container starts. They are actually like regular containers except they always run to completion and each initContainer should complete successfully before the next one starts. Refer https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#understanding-init-containers  

> NOTE: Grouping multiple co-located and co-managed containers in a single Pod is a relatively advanced use case. You should use this pattern only in specific instances in which your containers are tightly coupled.

## How to perform OS upgrade on a node in k8s cluster ?

- Sometimes OS upgrades may require restarts or can have longer down times for some OS related reasons. In that case we can perform the below steps. 
- Drain the node and then perform maintenance activity and get the node back up. Post that uncordon the node which will make it schedulable. Follow these steps
```
$ kubectl drain <node_name>
# You may have to ignore daemonsets
$ kubectl drain <node_name> --ignore-daemonsets
```
- perform maintenance activity. Post that when it is up.
```
$ kubectl uncordon <node_name>
```
- Before the node is drained, the command first cordons(makes it unschedulable) the node and then drains it. While draining the node it evicts the pod if they are managed by a ReplicationController, ReplicaSet, Deployment. An unmanaged pod will not be evicted by default. You will have to use force option in that case `kubectl drain <node_name> --ignore-daemonsets --force`  

## Cluster upgrades

- Cannot upgrade directly to new version. Have to upgrade minor to minor versions. For eg you cannot jump from 1.11 to 1.13. The cluster have to be upgraded to 1.12 first. 
- All core k8s components cannot be higher than kube-api server. Kubectl version can be in this range: api_server_version - 1  < kubectl_version > api_server_version + 1
- etcd's and coredns versioning is maintained separately.
- Overall process is to bring down the master node first. Upgrade master node, bring it back up and then upgrade the worker nodes

### Strategies to upgrade worker nodes

- bring all down and then bring them all back up 
- drain one, shift workload on other worker nodes and then upgrade this node. Upgrade one worker node at a time.
- Add same no of new worker nodes with updated versions of kubeadm,kubelet and kubectl and then drain the old ones.

## Backup and restore 

- `export ETCDCTL_API=3`
- `etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/snapshot-pre-boot.db`
- `etcdctl --endpoints=https://127.0.0.1:2379 snapshot restore /opt/snapshot-pre-boot.db --data-dir /var/lib/etcd/data-from-backup`
- Change the data-dir in manifest yaml in /etc/kubernetes/manifests/ directory.


