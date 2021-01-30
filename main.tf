resource "helm_release" "nats-operator" {
  name       = "nats-operator"
  repository = "https://nats-io.github.io/k8s/helm/charts/"
  chart      = "nats-operator"
  version    = "0.7.4"
  namespace  = "nats-io"

  values = [
    <<EOF
    cluster:
      create: true
      name: nats-cluster
      auth:
        enabled: false
      tls:
        enabled: false
    EOF
  ]

  create_namespace = true
}

resource "helm_release" "nats-streaming" {
  depends_on = [ helm_release.nats-operator ]

  name       = "nats-stan"
  repository = "https://nats-io.github.io/k8s/helm/charts/"
  chart      = "stan"
  version    = "0.7.4"
  namespace  = "nats-io"

  values = [
    <<EOF
    stan:
      clusterID: demo
      nats:
        url: nats://nats-cluster.nats-io.svc.cluster.local:4222
      logging:
        debug: true
        trace: true
    store:
      volume:
        storageClass: hostpath
    EOF
  ]
  
  create_namespace = true
}