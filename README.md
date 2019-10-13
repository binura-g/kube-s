KUBE-S
---

A lightweight CLI tool for quickly finding specific k8s resources (by pattern matching Names) across all clusters available to kubectl.

### Why not use a bash script? kube-s is **FAST**!

**kube-s searches through all your clusters concurrently and is much faster than searching through each cluster with something like `grep`.**
 
In general, `kube-s` outperforms an equivalent bash script search by a few good seconds. The higher the number of clusters, the more significant this difference becomes.
 
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

Install globally using go-get (Requires Go 1.13+)

`go get github.com/BinuraG/kube-s`

or Install from Release Build
 - Download the release specific to your OS from `./release`
 - Add the executable to your $PATH variable
