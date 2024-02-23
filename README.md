# cluster-controller

cluster-controller for Weave GitOps Enterprise.

This provides an in-cluster representation of clusters.

A `GitOpsCluster` can reference an existing CAPI cluster.

```yaml
apiVersion: gitops.weave.works/v1alpha1
kind: GitOpsCluster
metadata:
  name: dev
  namespace: default
spec:
  capiClusterRef:
    name: dev
```

Alternatively, a `GitOpsCluster` can reference a secret containing a kube config
file.

```yaml
apiVersion: gitops.weave.works/v1alpha1
kind: GitOpsCluster
metadata:
  name: dev
  namespace: default
spec:
  secretRef:
    name: dev
```

This code was originally written for Weave GitOps Enterprise.
