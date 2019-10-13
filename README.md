kube-s
---

A CLI tool for quickly finding specific k8s resources (by pattern matching Names) across all clusters available to kubectl.

### Why not use a bash script? `kube-s` is **FAST**

**kube-s searches through all the clusters concurrently and is much `fast`er than searching through each cluster with something like `grep`.**
 
In general, `kube-s` outperforms an equivalent bash script search by a few seconds. The higher the number of clusters, the more significant this difference becomes.
 
> You can find the scripts to run benchmarks under `./benchmark`.

## Usage

`$ kube-s <ResourceKind> <Pattern>`

Eg.
`$ kube-s pods my-`

Result: 
```
namespace-01 my-app-1
namespace-01 my-app-2
namespace-02 my-app-1
```

## Installation

