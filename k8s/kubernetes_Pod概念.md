# Pod
在Kubernetes中，Pod是最小的可部署计算单元，可以创建、管理和调度。一个Pod可以包含一个或多个容器，这些容器共享网络和存储资源。

Pod是一组（一个或多个） 容器； 这些容器共享存储、网络、以及怎样运行这些容器的声明。 Pod 中的内容总是并置（colocated）的并且一同调度，在共享的上下文中运行。 Pod 所建模的是特定于应用的 “逻辑主机”，其中包含一个或多个应用容器， 这些容器相对紧密地耦合在一起。 在非云环境中，在相同的物理机或虚拟机上运行的应用类似于在同一逻辑主机上运行的云应用。

除了应用容器，Pod 还可以包含在 Pod 启动期间运行的 Init 容器。 你也可以注入临时性容器来调试正在运行的 Pod。

# 什么是 Pod？
Pod 的共享上下文包括一组 Linux 名字空间、控制组（cgroup）和可能一些其他的隔离方面， 即用来隔离容器的技术。 在 Pod 的上下文中，每个独立的应用可能会进一步实施隔离。

Pod 类似于共享名字空间并共享文件系统卷的一组容器。

Kubernetes 集群中的 Pod 主要有两种用法：

- [运行单个容器的 Pod ] 每个 Pod 一个容器"模型是最常见的 Kubernetes 用例； 在这种情况下，可以将 Pod 看作单个容器的包装器，并且 Kubernetes 直接管理 Pod，而不是容器。
- [运行多个协同工作的容器的 Pod] Pod 可以封装由紧密耦合且需要共享资源的多个并置容器组成的应用。 这些位于同一位置的容器构成一个内聚单元。 将多个并置、同管的容器组织到一个 Pod 中是一种相对高级的使用场景。 只有在一些场景中，容器之间紧密关联时你才应该使用这种模式。

# 使用 Pod

下面是一个 Pod 示例，它由一个运行镜像 nginx:1.14.2 的容器组成。
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80
```

要创建上面显示的 Pod，将他保存为一个pod.yaml文件，然后运行以下命令：
```bash
kubectl apply -f pod.yaml
```

Pod 通常不是直接创建的，而是使用工作负载资源创建的。

# 用于管理 Pod 的工作负载资源
通常你不需要直接创建 Pod，甚至单实例 Pod。相反，你会使用诸如 Deployment 或 Job 这类工作负载资源来创建 Pod。 如果 Pod 需要跟踪状态，可以考虑 StatefulSet 资源。

每个 Pod 都旨在运行给定应用程序的单个实例。如果希望横向扩展应用程序 （例如，运行多个实例以提供更多的资源），则应该使用多个 Pod，每个实例使用一个 Pod。 在 Kubernetes 中，这通常被称为副本（Replication）。 通常使用一种工作负载资源及其控制器来创建和管理一组 Pod 副本。

# Pod 模板
工作负载资源的控制器通常使用 Pod 模板（Pod Template）来替你创建 Pod 并管理它们。

Pod 模板是包含在工作负载对象中的规范，用来创建 Pod。这类负载资源包括 Deployment、 Job 和 DaemonSet 等。

工作负载的控制器会使用负载对象中的 PodTemplate 来生成实际的 Pod。 PodTemplate 是你用来运行应用时指定的负载资源的目标状态的一部分。

创建 Pod 时，你可以在 Pod 模板中包含 Pod 中运行的容器的环境变量。

下面的示例是一个简单的 Job 的清单，其中的 template 指示启动一个容器。 该 Pod 中的容器会打印一条消息之后暂停。

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: hello
spec:
  template:
    # 这里是 Pod 模板
    spec:
      containers:
      - name: hello
        image: busybox:1.28
        command: ['sh', '-c', 'echo "Hello, Kubernetes!" && sleep 3600']
      restartPolicy: OnFailure
    # 以上为 Pod 模板
```

修改 Pod 模板或者切换到新的 Pod 模板都不会对已经存在的 Pod 直接起作用。 如果改变工作负载资源的 Pod 模板，工作负载资源需要使用更新后的模板来创建 Pod， 并使用新创建的 Pod 替换旧的 Pod。

例如，StatefulSet 控制器针对每个 StatefulSet 对象确保运行中的 Pod 与当前的 Pod 模板匹配。如果编辑 StatefulSet 以更改其 Pod 模板， StatefulSet 将开始基于更新后的模板创建新的 Pod。

每个工作负载资源都实现了自己的规则，用来处理对 Pod 模板的更新。

在节点上，kubelet 并不直接监测或管理与 Pod 模板相关的细节或模板的更新，这些细节都被抽象出来。 这种抽象和关注点分离简化了整个系统的语义， 并且使得用户可以在不改变现有代码的前提下就能扩展集群的行为。