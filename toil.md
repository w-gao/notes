# toil

## Clusters

#### Launch a mesos (or kubernetes) static cluster

- all nodes with the given node types will be created by the `toil launch-cluster` command

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

* make sure to specify the zone

```shell
toil destroy-cluster -z us-west-2a **cluster-name** --logDebug
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

