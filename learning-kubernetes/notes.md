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
- Helpful commands
```
kubectl rollout status deployment/<deployment_name>
kubectl rollout history deployment/<deployment_name>
kubectl rollout undo deployment/<deployment_name>
```

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

## controller manager

- brain of monitoring inside kubernetes.
- custom options like "--pod-eviction-timeout,--node-monitor-period,etc" are available in the manifest file.

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

## Monitoring 

- Earlier heapster, got depracated. Now Metrics Server but it is "in memory".
- You will have to deploy the metrics server explicitly 
- Commands to check utilisation
```
kubectl top node
kubectl top pod
```

## ConfigMaps

> For plugging in the entire configMap as env variables use below yaml config:
```
envFrom:
  - configMapRef:
      name: webapp-config-map
```

## Secrets

- When a secret is plugged in as a volume, every key becomes a file on the path you mount the volume on. 
- You can plug in the secret as environment variables also. 
- They are in encoded(base64) format inside the `data:` field. 


## Multiple containers

- A pod can encapsulate an application composed of multiple co-located containers that are tightly coupled and need to share resources. 
- Pods natively provide 2 kinds of shared resources for their containers: 
    - network
    - storage
- These co-located containers form a single cohesive unit of service—for example, one container serving data stored in a shared volume to the public, while a separate sidecar container refreshes or updates those files. The Pod wraps these containers, storage resources, and an ephemeral network identity together as a single unit.
- Refer https://kubernetes.io/docs/concepts/workloads/pods/#how-pods-manage-multiple-containers
- Pods also have initContainers. As the name suggests these are run before the main container starts. They are actually like regular containers except they always run to completion and each initContainer should complete successfully before the next one starts. Refer https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#understanding-init-containers  
- Also, init containers do not support lifecycle, livenessProbe, readinessProbe, or startupProbe because they must run to completion before the Pod can be ready

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

- Cannot upgrade directly to new version. Have to upgrade minor to minor versions. For eg you cannot jump from 1.11 to 1.13. The cluster has to be upgraded to 1.12 first. 
- All core k8s components cannot be higher than kube-api server. Kubectl version can be in this range: api_server_version - 1  < kubectl_version > api_server_version + 1
- etcd's and coredns versioning is maintained separately.
- Overall process is to bring down the master node first. Upgrade master node, bring it back up and then upgrade the worker nodes

### Strategies to upgrade worker nodes

- bring all down and then bring them all back up 
- drain one, shift workload on other worker nodes and then upgrade this node. Upgrade one worker node at a time.
- Add same no of new worker nodes with updated versions of kubeadm,kubelet and kubectl and then drain the old ones.

## Backup and restore 

- Check what version is being used first. If it is 2 the use api version 3 by setting this env variable. `export ETCDCTL_API=3`
- `etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/snapshot-pre-boot.db`
- `etcdctl --endpoints=https://127.0.0.1:2379 snapshot restore /opt/snapshot-pre-boot.db --data-dir /var/lib/etcd/data-from-backup`
- Change the data-dir in manifest yaml in /etc/kubernetes/manifests/ directory.

## Certificates API

- If a new user wants to access the k8s cluster but you dont want to give access of the master node or the kube-admin user, then this user can generate a csr.
- This csr can then be approved by the k8s certificates API. This is handled by the controller-manager. 
- By using kubectl you can create the CSR object and then approve it. 
    - `kubectl get csr myuser`
    - `kubectl certificate approve/deny myuser`

## kubeconfig

- Default location: `~/.kube/config`
- This also shows which users have access to the cluster. 
- The same kubeconfig can be used even when you have multiple k8s cluster. 
- Structure is divided into 3 main components:
    - clusters
    - contexts
    - users
- `context` ties the user with the cluster it needs access to. At one time, there can only be one context. 
- Below is what the defaul kubeconfig looks like. It has only one user ie kubernetes-admin
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://master_ip:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: REDACTED
    client-key-data: REDACTED
```

## Authorization

- The Kubernetes authorization server may authorize a request using one of the authorization modes listed below
  - Node -> A component will have to be in a system:nodes group to be authorized by this mode (eg: kubelet)
  - ABAC -> Attribute-based access control (Not very modular)
  - RBAC -> Role-based acccess control (quite similar to IAM roles in AWS, although represented in a different way)
  - Webhook -> Authorization handled by a 3rd party application
- Check in the kube-api-server manifest what mode of authorization is being used. A cluster made with `kubeadm` has `--authorization-mode=Node,RBAC` by default. Every request is authorised in the order from left to right if the former denies or doesnt support the authorication type.
- Every role and rolebinding created is restricted to it's namespace. You may have to create roles with same level of access for different namespaces. 
- Full documentation : https://kubernetes.io/docs/reference/access-authn-authz/authorization/

### RBAC

- RBAC is the recommended way to manage authorization in k8s. Below components make up RBAC.
    - Role
    - RoleBinding
    - ClusterRole
    - ClusterRoleBinding
- Roles are used to restrict access to namespaces. ClusterRoles are not namespaced, level of access granted in this case is across the cluster. 
- To assign a role to a user/group you can use rolebinding.
- To assign a clusterRole to a user/group you can use clusterrolebinding.
- Command to list namespace objects  `kubectl api-resources --namespaced=true`
- Command to list objects that are not namespaced `kubectl api-resources --namespaced=false`

### ServiceAccount

- Users can use certificates and rolebindings to get access to the Kubernetes API, but what about applications/bots. This is solved by serviceaccounts
- By default a pod will have the default serviceAccount configured. ServiceAccount gives a token which can be passed in the `Authorization` header. This token is available as a volumeMount by default at location /var/run/secrets/kubernetes.io/serviceaccount/token. 
- For your bot/application you can create a custom serviceaccount and plug it to the desired pod by using the `serviceAccountName` spec. 
- This custom servericeAccount would have to have a rolebinding in order to get access to the k8s resources.

## Network policy

- By default in k8s, all pods can access each other on all ports. The defaul network policy is "Allow all".
- But in a large cluster where several teams have their own applications it's good to have a network policy to restrict access to critical pods like a database. 
- Specs defined in the network policy in the list format are evaluated as an OR operation. Multiple specs under one spec are evaluated as an AND operation. 
> Refer : https://kubernetes.io/docs/concepts/services-networking/network-policies/#networkpolicy-resource

## Storgage : 

## Container Storage interface : 

- It is a standard for exposing arbitrary block and file storage systems to containerized workloads on Container Orchestration Systems(COs) like kubernetes.
- Using CSI third-party storage providers can write and deploy plugins exposing new storage systems in Kubernetes without ever having to touch the core Kubernetes code.

## Volumes :

- Pods have ephemeral storage by default. Data does not persist across pod restarts or failures.
- Volumes can help with data persistence. Kubernetes supports many types of volumes. A Pod can use any number of volume types simultaneously. Ephemeral volume types have a lifetime of a pod, but persistent volumes exist beyond the lifetime of a pod. When a pod ceases to exist, Kubernetes destroys ephemeral volumes; however, Kubernetes does not destroy persistent volumes. For any kind of volume in a given pod, data is preserved across container restarts.
- At its core, a volume is a directory, possibly with some data in it, which is accessible to the containers in a pod. How that directory comes to be, the medium that backs it, and the contents of it are determined by the particular volume type used.
> Refer : https://kubernetes.io/docs/concepts/storage/volumes/

### PersistentVolumes:

- It is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using storage classes. It is resource in a cluster just like how a node is a cluster resource.
- They are volume plugins like Volumes, but their lifecycle is independent of any individual pod that uses the PV. 
> Refer : https://kubernetes.io/docs/concepts/storage/persistent-volumes/

### PersistentVolumeClaims:

- It is a request for storage by a user. The way pods consume node resources, persistentvolumeclaims consume PV resources. There has to be a 1:1 mapping for pvc and pv. 
- Requests by pvc should match the stuff that the pv has provisioned. If the capacity requested is lower than the pv available and no other suitable pv is available then the available pv is selected. The remaining capacity is not given to any other pvc.
> Refer : https://kubernetes.io/docs/concepts/storage/persistent-volumes/

### StorageClasses:

- Usually an administrator provisons a persistent volume and then creates a PVC. After the PVC is in a bound state only then the Pod can use it. This is called satic provisioning. 
- In real world scenarious, this can become very tedious. For this dynamic provisioning has to be used. StorageClasses can be used to solve this.
- It removes the need for persistent volumes. You can directly write a PVC using a storage class.
> Refer : https://kubernetes.io/docs/concepts/storage/storage-classes/

## Networking in k8s 

- The creation of network interfaces, attachment and ip address management is handled by the CNI plugin.
- Every pod launched gets a unique ip address defined in the ip address range setting mentioned with the CNI plugin.
- Run `kubectl get pods -o wide` to check the ip address assigned to every pod. 
- If you ever notice the cluster ip addresses assigned to services in k8s, the range of these ip addresses are different than the pod's ip addresses. These IP addresses are virtual and come from a pool of IP range mentioned in api-server manifest file. 
- So how do pods connect to services internally if the IP is virtual? 
- It actually goes via kube-proxy. kube-proxy itself has a node's ip address. kube-proxy by default uses the iptables mode which sets up NAT. All these rules are updated with every deployment. 
- Run `sudo iptables -nvL -t nat` to check NAT rules created for k8s services.
> Refer : https://kubernetes.io/docs/concepts/services-networking/service/#configuration

## JSON path

Tutorial : https://www.youtube.com/watch?v=XHRsPp6TORo

- kubectl get commands convert the long json output into a short readable format. Sometimes to get internal details we can use kubectl get nodes -o=jsonpath='{query}'

## Kubernetes admission controllers

- https://www.youtube.com/watch?v=P7QAfjdbogY

## Pod security policy

- built-in admission controller that allows a cluster administrator to control security-sensitive aspects of the Pod specification.
- not to be confused with pod security context (https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)



