#!/bin/bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Spider

set -o errexit -o nounset -o pipefail -o xtrace

CURRENT_FILENAME=$( basename $0 )

[ -z "${HTTP_PROXY}" ] || export https_proxy=${HTTP_PROXY}

[ -z "$E2E_CLUSTER_NAME" ] && echo "error, miss E2E_CLUSTER_NAME " && exit 1
echo "$CURRENT_FILENAME : E2E_CLUSTER_NAME $E2E_CLUSTER_NAME "

[ -z "$E2E_IP_FAMILY" ] && echo "error, miss E2E_IP_FAMILY" && exit 1
echo "$CURRENT_FILENAME : E2E_IP_FAMILY $E2E_IP_FAMILY "

[ -z "$METALLB_VERSION" ] && echo "error, miss METALLB_VERSION" && exit 1
echo "$METALLB_VERSION : METALLB_VERSION $METALLB_VERSION "

[ -z "$E2E_KUBECONFIG" ] && echo "error, miss E2E_KUBECONFIG " && exit 1
[ ! -f "$E2E_KUBECONFIG" ] && echo "error, could not find file $E2E_KUBECONFIG " && exit 1
echo "$CURRENT_FILENAME : E2E_KUBECONFIG $E2E_KUBECONFIG "

helm repo remove metallb &>/dev/null || true
helm repo add metallb https://metallb.github.io/metallb
helm repo update metallb

HELM_OPTIONS="	--set controller.image.repository=${E2E_METALLB_IMAGE_REPO}/metallb/controller \
                --set speaker.image.repository=${E2E_METALLB_IMAGE_REPO}/metallb/speaker \
                --set speaker.frr.image.repository=${E2E_METALLB_IMAGE_REPO}/frrouting/frr  --set speaker.frr.enabled=true " ; \

IMAGE_LIST=` helm template --version=${METALLB_VERSION}  metallb/metallb ${HELM_OPTIONS} | grep ' image: ' | tr -d '"' | awk '{print $2}'  | sort | uniq | tr '\n' ' '  ` ; \

if [ -z "${IMAGE_LIST}" ] ; then \
	echo "warning, failed to find image from chart template" ; \
else \
	echo "found image from chart template: ${IMAGE_LIST} " ; \
	for IMAGE in ${IMAGE_LIST} ; do \
		EXIST=` docker images | awk '{printf("%s:%s\n",$1,$2)}' | grep "${IMAGE}" ` || true ; \
		if [ -z "${EXIST}" ] ; then \
			echo "docker pull ${IMAGE} to local" ; \
			docker pull ${IMAGE} ;\
		fi ;\
		echo "load local image ${IMAGE} " ; \
		kind load docker-image ${IMAGE}  --name ${E2E_CLUSTER_NAME} ; \
	done ; \
fi ; \

helm upgrade --install -n kube-system metallb  metallb/metallb --kubeconfig=${E2E_KUBECONFIG} \
    --version=${METALLB_VERSION} --wait --debug --timeout 5m \
	--set speaker.nodeSelector."kubernetes\.io/hostname"=${E2E_CLUSTER_NAME}-control-plane \
	${HELM_OPTIONS}

sleep 5
echo "wait metallb related pod running ..."
kubectl wait --for=condition=ready -l app.kubernetes.io/instance=metallb --timeout=300s pod -n kube-system \
    --kubeconfig ${E2E_KUBECONFIG} 

METALLB_CR_TEMPLATE='
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: default-pool
  namespace: kube-system
spec:
  autoAssign: true
  addresses:
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: default-arp
  namespace: kube-system
'

echo "apply ippool"
echo "docker network inspect kind"
docker network inspect kind
case ${E2E_IP_FAMILY} in
  ipv4)
    Subnet1=$(docker network inspect kind -f {{\(index\ $.IPAM.Config\ 0\).Subnet}}) ; \
    Subnet1=${Subnet1%%/*} ; \
    if ! grep -iE "[a-f]" <<< "${Subnet1}" ; then \
        IPPOO1="${Subnet1%0}" ; \
        IPPOO1="${IPPOO1}50-${IPPOO1}90" ; \
    else \
        Subnet2=$(docker network inspect kind -f {{\(index\ $.IPAM.Config\ 1\).Subnet}}) ; \
        Subnet2=${Subnet2%%/*} ; \
        if ! grep -iE "[a-f]" <<< "${Subnet2}" ; then \
            IPPOO1="${Subnet2%0}" ; \
            IPPOO1="${IPPOO1}50-${IPPOO1}90" ; \
        else \
            echo "failed to find node ipv4 subnet" ; \
            exit 1 ; \
        fi ; \
    fi ; \
    echo "IPPOO1: ${IPPOO1}" ; \
    echo "${METALLB_CR_TEMPLATE}" \
        | sed '/addresses:/ a\    - '"${IPPOO1}"''  \
        | kubectl --kubeconfig=${E2E_KUBECONFIG} apply -f - ; \
    ;;
  ipv6)
    Subnet1=$(docker network inspect kind -f {{\(index\ $.IPAM.Config\ 0\).Subnet}}) ; \
    Subnet1=${Subnet1%%/*} ; \
    if grep -iE "[a-f]" <<< "${Subnet1}" ; then \
        IPPOO1="${Subnet1}50-${Subnet1}90" ; \
    else \
        Subnet2=$(docker network inspect kind -f {{\(index\ $.IPAM.Config\ 1\).Subnet}}) ; \
        Subnet2=${Subnet2%%/*} ; \
        if grep -iE "[a-f]" <<< "${Subnet2}" ; then \
            IPPOO1="${Subnet2}50-${Subnet2}90" ; \
        else \
            echo "failed to find node ipv6 subnet" ; \
            exit 1 ; \
        fi ; \
    fi ; \
    echo "IPPOO1: ${IPPOO1}" ; \
    echo "${METALLB_CR_TEMPLATE}" \
        | sed '/addresses:/ a\    - '"${IPPOO1}"''  \
        | kubectl --kubeconfig=${E2E_KUBECONFIG} apply -f - ; \
    ;;
  dual)
    Subnet1=$(docker network inspect kind -f {{\(index\ $.IPAM.Config\ 0\).Subnet}}) ; \
    Subnet1=${Subnet1%%/*} ; \
    if grep -iE "[a-f]" <<< "${Subnet1}" ; then \
        IPPOO1="${Subnet1}50-${Subnet1}90" ; \
    else \
        IPPOO1="${Subnet1%0}" ; \
        IPPOO1="${IPPOO1}50-${IPPOO1}90" ; \
    fi ; \
    Subnet2=$(docker network inspect kind -f {{\(index\ $.IPAM.Config\ 1\).Subnet}}) ; \
    Subnet2=${Subnet2%%/*} ; \
    if grep -iE "[a-f]" <<< "${Subnet2}" ; then \
        IPPOO2="${Subnet2}50-${Subnet2}90" ; \
    else \
        IPPOO2="${Subnet2%0}" ; \
        IPPOO2="${IPPOO2}50-${IPPOO2}90" ; \
    fi ; \
    echo "IPPOO1: ${IPPOO1}" ; \
    echo "IPPOO2: ${IPPOO2}" ; \
    echo "${METALLB_CR_TEMPLATE}" \
        | sed '/addresses:/ a\    - '"${IPPOO1}"''  \
        | sed '/addresses:/ a\    - '"${IPPOO2}"'' ; \
    echo "${METALLB_CR_TEMPLATE}" \
        | sed '/addresses:/ a\    - '"${IPPOO1}"''  \
        | sed '/addresses:/ a\    - '"${IPPOO2}"''  \
        | kubectl --kubeconfig=${E2E_KUBECONFIG} apply -f - ; \
    ;;
  *)
    echo "the value of E2E_IP_FAMILY: ipv4 or ipv6 or dual"
    exit 1
esac

sleep 1

echo -e "\033[35m Succeed to install metallb \033[0m"