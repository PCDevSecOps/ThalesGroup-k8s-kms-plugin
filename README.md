# K8S-KMS-Plugin

This microservice implements the Kubernetes KMS protocol as a gRPC service that leverages a remote or local HSM via PKCS11.

This plugin will also run in proxy mode which can connect to a remote plugin service running in a secure network device (Key Managers)

## requirements

This service is designed for kubernetes clusters that are using version 1.10.0 or higher and implements the KMS API:

https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/

So for development purposes, you'll want a cluster that can be configured to use a KMS gRPC endpoint on your APIServer nodes. 

Locally you should install [skaffold.dev](https://skaffold.dev) tooling as well as Cloud Code in your favorite IDE to leverage the skaffold.yaml file in this repo.

## Quick Start for testing

- Get a k8s cluster to deploy the `k8s-kms-plugin`
- Install skaffold.dev  [skaffold.dev](https://skaffold.dev) 
- Running `skaffold dev` or `make dev` should put the  stack into a local deployment pipeline of the plugin being tested as a KMS gRPC server.

## Deployment scenarios

This plugin is designed to be deployed in 2 configurations

- Client/Server - k8s-kms-plugin in `client` mode will `enroll` to an external k8s-kms-plugin running in `serve` mode
- StandAlone(TODO) - Plugin and PKCS11 library deployed as StaticPod/HostContainer on APIServer nodes, this will require
coordination with k8s provisioning tools.

## Development Environment

`k8s` houses some sample client and server deployments for e2e testing until such time as this plugin is 100% network functional,
 and we can move it to a CICD pattern, as we'll have many actors to coordinate. 

All apis are defined in the `/apis` dir, and as we iterate on the spec docs, one must then run `make gen` and refactor 
until the 2 stacks come up

Both EST and KMS-Plugin binaries are in the `/cmd` dir
 
The `Makefile` contains commands for easy execution:
- `make gen` - generates all apis into gRPC or OpenAPI Servers and Clients
- `make dev` - loads project into your kubernetes cluster (minikube or GKE will work just fine), and continously builds and deploys as you develop.
- `make build` - builds the standalone `k8s-kms-plugin` binary

NOTE:  Currently the standalone plugin just waits for the 

## Debug Environment

For a remote debug, build the plugin with debug mode :

```sh
go get github.com/go-delve/delve/cmd/dlv
make build-debug
```

### TPM2

This documentation is based on a Debian12 virtual machine experiment and a software TPM.

1. Install a software tpm, its dependencies
2. Install the libraries `libtpms`, `libtpm2-pkcs11` and `tpm2-abrmd`
3. Add your user to the group `tss` and launch the software tpm
4. Initialize the PKCS11 store, tokens and a user pin :

```sh
export STORE="$HOME/tpm2"
export MODULE="/usr/lib/x86_64-linux-gnu/libtpm2_pkcs11.so.1"
# store
mkdir -p $STORE
tpm2_ptool init --path=$STORE
# token
pkcs11-tool --module $MODULE --slot-index 0 --init-token --label mylabel --so-pin mysopin
# an initialized token should appears :
pkcs11-tool --module $MODULE --list-token-slots
# pin
pkcs11-tool --module $MODULE --init-pin --so-pin mysopin --login --pin mypin --slot-index 0
# test
pkcs11-tool --module $MODULE --label mylabel --pin mypin --generate-random 16 | xxd
```

5. Create the required keys :

```sh
# aes
tpm2_ptool addkey --algorithm aes256 --label mylabel --key-label aes0 --userpin mypin --path $STORE
# hmac
tpm2_ptool addkey --algorithm hmac:sha256 --label mylabel --key-label hmac0 --userpin mypin --path $STORE
# list
pkcs11-tool --module $MODULE --token-label mylabel --pin mypin -O
tpm2_ptool listobjects --label mylabel --path $STORE
```

7. Create the socket's folder for the kms. The socket will be automatically created inside.  
**WARNING: do not use `/tmp` for production environment !**

```sh
mkdir -p /tmp/run
```

8. Launch the plugin :

```sh
# local debug
k8s-kms-plugin serve \
  --provider p11 --p11-lib $MODULE --p11-key-label aes0 --p11-label mylabel --p11-pin mypin --algorithm aes-cbc --enable-server
# remote debug
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient --check-go-version=false \
  exec k8s-kms-plugin-dlv -- serve \
  --provider p11 --p11-lib $MODULE --p11-key-label aes0 --p11-label mylabel --p11-pin mypin --algorithm aes-cbc --enable-server
```
