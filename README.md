KUBE-S
---

[![license](https://img.shields.io/github/license/DAVFoundation/captain-n3m0.svg?style=flat-square)](https://github.com/binura-g/kube-s/blob/master/LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

A lightweight CLI tool for quickly finding specific k8s resources (by pattern matching Names) across all clusters available to kubectl.

### Why not use a bash script? kube-s is **FAST**!

**kube-s searches through all your clusters concurrently and is much faster than searching through each cluster with something like `grep`.**
 
In general, `kube-s` outperforms an equivalent bash script search by a few good seconds. The higher the number of clusters, the more significant this difference becomes.
 
> You can find the scripts to run benchmarks under `./benchmark`.

## Usage

`$ kube-s <ResourceKind> <Pattern>`

*Eg.* Search for all pods with names matching `"my-"`

`$ kube-s pods my-`

Result: 
```
cluster-01    namespace-01    my-app-1
cluster-02    namespace-01    my-app-2
cluster-03    namespace-02    my-app-1
```

> kube-s searches all clusters available in your kubeconfig

## Installation

Install globally using go-get (Requires Go 1.13+)

`go get github.com/binura-g/kube-s`

or Install from Release Build
 - Download the release specific to your OS from `./release`
 - Add the executable to your $PATH

