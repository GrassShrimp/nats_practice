# NATS Practice
This is a demo for test nats queue

## Prerequisites
- [terraform](https://www.terraform.io/downloads.html)
- [docker](https://www.docker.com/products/docker-desktop) and enable kubernetes
- [skaffold](https://skaffold.dev/docs/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize)

## Usage
check current context of kubernetes is __docker-desktop__
```bash
$ kubectl config current-context
```
initialize terrafrom module
```bash
$ terraform init
```
launch nats and nats streaming
```bash
$ terraform apply -auto-approve
```
for test nats publish and subscribe
```bash
$ cd nats-publish && skaffold dev
$ cd nats-subscribe && skaffold dev
```
for test nats request and reply
```bash
$ cd nats-request && skaffold dev
$ cd nats-reply && skaffold dev
```
for test nats-streaming publish and subscribe
```bash
$ cd stan-publish && skaffold dev
$ cd stan-subscribe && skaffold dev
```