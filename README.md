# Go Dumb CRI

The dumb container runtime trying to be compatible with Kubernetes CRI.

## Usage

Run the server and create an IPC socket in `/tmp/go-dumbcri.sock`

```bash
go run .
```

Install `crictl` and use it as the client

```bash
brew install cri-tools
crictl --runtime-endpoint unix:///tmp/go-dumbcri.sock version
```

The output should look like this

```text
Version:  v1alpha2
RuntimeName:  go-dumbcri
RuntimeVersion:  v0.0.0-alpha.1
RuntimeApiVersion:  v0.0.0-alpha.1
```

## Why

`go-dumbcri` is a small project to learn about the internals of Kubernetes CRI.

## Lessons Learned

What I have understood so far is that the container runtime interface is defined in protobuf in <https://github.com/kubernetes/kubernetes/tree/v1.22.3/staging/src/k8s.io/cri-api>.

The code is extracted into a dedicated repository at <https://github.com/kubernetes/cri-api/tree/v0.22.3>. Or maybe the other way around...

It is then imported to implement either a client ( <https://github.com/kubernetes-sigs/cri-tools/blob/ed2adf0/go.mod#L27> ) or a server ( <https://github.com/containerd/containerd/blob/d418660/go.mod#L75>, <https://github.com/cri-o/cri-o/blob/3943334/go.mod#L70> ).

## Documentation

- <https://github.com/kubernetes/cri-api>
- <https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/>
- <https://kubernetes.io/docs/setup/production-environment/container-runtimes/>
