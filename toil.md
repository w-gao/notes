# toil

## Clusters

#### Launch a mesos (or kubernetes) static cluster

- The nodes with the given node types will be created at once when the `toil launch-cluster` command is run.

```shell
TOIL_APPLIANCE_SELF=quay.io/ucsc_cgl/toil:5.3.0 \
toil launch-cluster **cluster-name** \
      --leaderNodeType t2.small \
      --zone us-west-2a \
      --keyPairName **** \
      --nodeTypes t2.small,t2.micro \
      --workers 1,1 \
      --logDebug
```

#### Launch a mesos (or kubernetes) auto scaling cluster

```shell
TOIL_APPLIANCE_SELF=quay.io/ucsc_cgl/toil:5.3.0 \
toil launch-cluster **cluster-name** \
      --provisioner aws \
      --zone us-west-2a \
      --leaderNodeType t2.medium \
      --leaderStorage 50 \
      --keyPairName **** \
      --nodeTypes t2.small \
      --nodeStorage 20 \
      --workers 0-1 \
      --clusterType kubernetes \
      --logDebug
```

#### Destroy a cluster

- make sure to specify the zone

```shell
toil destroy-cluster -z us-west-2a **cluster-name** --logDebug
```


## The Kubernetes batch system

For running Toil on Kubernetes, the docs is a great starting point: https://toil.readthedocs.io/en/latest/running/cloud/kubernetes.html.
Here are the notes that I jot down when following this tutorial.


#### Setting up a local cluster

Setting up a local k8s cluster was easy enough, with a few caveats:


1. Configuring AWS access key as a kubectl secret



2. Enabling the metrics server

Toil uses the metrics server to check if the pods are stuck in OOM [here](https://github.com/DataBiosphere/toil/blob/releases/5.7.x/src/toil/batchSystems/kubernetes.py#L760-L762).
It is probably already enabled on the GI cluster, but not on a minikube cluster by default.

Enabling it was simple, though, with this command:

```shell
minikube addons enable metrics-server
```

See: https://kubernetes.io/docs/tutorials/hello-minikube/#enable-addons



#### minikube dashboard


Don't forget to use the minikube dashboard, which provides a great UI for what's going on in the cluster.

```shell
minikube dashboard
```


## EC2

* commands to check logs on EC2
```
journalctl -b -t ignition --no-pager
journalctl -xeu kubelet

systemctl status toil-leader.service
systemctl status toil-worker.service

systemctl status create-kubernetes-cluster.service
```

