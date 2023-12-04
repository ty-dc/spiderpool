# Kind Quick Start

[**English**](./get-started-kind.md) | **简体中文**

Kind 是一个使用 Docker 容器节点运行本地 Kubernetes 集群的工具。Spiderpool 提供了安装 Kind 集群的脚本，您可以使用它来部署符合您需求的集群，进行 Spiderpool 的测试与体验。

## 先决条件

* 获取 Spiderpool 稳定版本的代码到本地主机上，并进入 Spiderpool 工程的根目录。

    ```bash
    ~# LATEST_RELEASE_VERISON=$(curl -s https://api.github.com/repos/spidernet-io/spiderpool/releases | grep '"tag_name":' | grep -v rc | grep -Eo "(v[0-9]+\.[0-9]+\.[0-9])" | sort -r | head -n 1)
    ~# curl -Lo /tmp/$LATEST_RELEASE_VERISON.tar.gz https://github.com/spidernet-io/spiderpool/archive/refs/tags/$LATEST_RELEASE_VERISON.tar.gz
    ~# mkdir -p /tmp/spiderpool && tar -xvf /tmp/$LATEST_RELEASE_VERISON.tar.gz -C /tmp/spiderpool
    ~# cd /tmp/spiderpool
    ```
  
* 执行 `make dev-doctor`，检查本地主机上的开发工具是否满足部署 Kind 集群与 Spiderpool 的条件。

    构建 Spiderpool 环境需要具备 Kubectl、Kind、Docker、Helm、yq 工具。如果你的本机上缺少，请运行 `test/scripts/install-tools.sh` 来安装它们。

## 快速启动

=== "创建基于 Spiderpool 单 CNI 环境"

    如下命令将创建一个 Macvlan 的单 CNI 网络环境。

    ```bash
    ~# make setup_singleCni_macvlan
    ```

=== "创建基于 Spiderpool 和 Calico 的双 CNI 环境"

    如下命令将创建一个 Calico 为 main CNI 并搭配 Macvlan 的多 CNI 网络环境。

    ```bash
    ~# make setup_dualCni_calico
    ```

=== "创建基于 Spiderpool 和 Cilium 的双 CNI 环境"

    如下命令将创建一个 Cilium 为 main CNI 并搭配 Macvlan 的多 CNI 网络环境。

    ```bash
    ~# make setup_dualCni_cilium
    ```

    > 确认操作系统 Kernel 版本号是是否 >= 4.9.17，内核过低时将会导致安装失败，推荐 Kernel 5.10+ 。

如果您在中国大陆，安装时可以额外指定参数 `-e E2E_CHINA_IMAGE_REGISTRY=true` ，以帮助您更快的拉取镜像。

## 验证安装

在 Spiderpool 工程的根目录下执行如下命令，为 kubectl 配置 Kind 集群的 KUBECONFIG。

```bash
~# export KUBECONFIG=$(pwd)/test/.cluster/spider/.kube/config
```

您可以看到类似如下的内容输出：

```bash
~# kubectl get nodes
NAME                   STATUS   ROLES           AGE     VERSION
spider-control-plane   Ready    control-plane   2m29s   v1.26.2
spider-worker          Ready    <none>          2m58s   v1.26.2

~# kubectll get po -n kube-sysem | grep spiderpool
NAME                                           READY   STATUS      RESTARTS   AGE                                
spiderpool-agent-4dr97                         1/1     Running     0          3m
spiderpool-agent-4fkm4                         1/1     Running     0          3m
spiderpool-controller-7864477fc7-c5dk4         1/1     Running     0          3m
spiderpool-controller-7864477fc7-wpgjn         1/1     Running     0          3m
spiderpool-init                                0/1     Completed   0          3m

~# kubectl get spiderippool
NAME                VERSION   SUBNET                    ALLOCATED-IP-COUNT   TOTAL-IP-COUNT   DEFAULT
default-v4-ippool   4         172.18.0.0/16             5                    253              true      
default-v6-ippool   6         fc00:f853:ccd:e793::/64   5                    253              true      
vlan100-v4          4         172.100.0.0/16            0                    2559             false
vlan100-v6          6         fd00:172:100::/64         0                    65009            false
vlan100-v4          4         172.200.0.0/16            0                    2559             false
vlan200-v6          6         fd00:172:200::/64         0                    65009            false
```

Spiderpool 提供的快速安装 Kind 集群脚本会自动为您创建一个应用，以验证您的 Kind 集群是否能够正常工作，以下是应用的运行状态：

```bash
~# kubectl get po -l app=test-pod -o wide
NAME                       READY   STATUS    RESTARTS   AGE     IP             NODE            NOMINATED NODE   READINESS GATES
test-pod-856f9689d-876nm   1/1     Running   0          5m34s   172.18.40.63   spider-worker   <none>           <none>
```

您也可以手动创建应用验证 Kind 集群是否能够正常工作，以下命令会创建 1 个副本 Deployment：

```bash
cat <<EOF | kubectl create -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  template:
    metadata:
      labels:
        app: test-app
    spec:
      containers:
      - name: test-app
        image: nginx
        imagePullPolicy: IfNotPresent
        ports:
        - name: http
          containerPort: 80
          protocol: TCP
EOF
```

```bash
~# kubectl get po -l app=test-app -o wide
NAME                        READY   STATUS    RESTARTS   AGE     IP              NODE                   NOMINATED NODE   READINESS GATES
test-app-84d5699474-dbtl5   1/1     Running   0          6m23s   172.18.40.112   spider-control-plane   <none>           <none>
```

通过测试，Kind 集群一切正常，您可以基于它测试与体验 Spiderpool 的更多功能。

## 卸载

* 卸载 Kind 集群

    执行 `make clean` 卸载 Kind 集群。

* 删除测试镜像

  ```bash
  ~# docker rmi -f $(docker images | grep spiderpool | awk '{print $3}')
  ~# docker rmi -f $(docker images | grep multus | awk '{print $3}')
  ```
