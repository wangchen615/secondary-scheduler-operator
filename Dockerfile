FROM registry.svc.ci.openshift.org/openshift/release:golang-1.16 AS builder
WORKDIR /go/src/github.com/openshift/secondary-scheduler-operator
COPY . .
# image-references file is not recognized by OLM
RUN rm manifests/4*/image-references
RUN go build -o secondary-scheduler-operator ./cmd/cluster-kube-descheduler-operator

FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base
COPY --from=builder /go/src/github.com/openshift/secondary-scheduler-operator/secondary-scheduler-operator /
# Upstream bundle and index images does not support versioning so
# we need to copy a specific version under /manifests layout directly
COPY --from=builder /go/src/github.com/openshift/secondary-scheduler-operator/manifests/4.8/* /manifests/
COPY --from=builder /go/src/github.com/openshift/secondary-scheduler-operator/metadata /metadata

ENTRYPOINT ["/secondary-scheduler-operator"]
LABEL io.k8s.display-name="OpenShift Secondary-scheduler Operator" \
      io.k8s.description="This is a component of OpenShift and manages the secondary scheduler" \
      io.openshift.tags="openshift,secondary-scheduler-operator" \
      com.redhat.delivery.appregistry=true \
      maintainer="AOS workloads team, <aos-workloads@redhat.com>"
