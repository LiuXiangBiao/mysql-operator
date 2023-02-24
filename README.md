# mysql-operator

## 介绍和功能

- 基于operator的k8s开发，做的有状态服务statefulset部署mysql实现数据
持久化，会自动创建及更新：pvc、pv、configmap、secret、service、statefulset 资源
  (使用的nfs底层供应存储)
- configmap 存储数据库配置文件信息挂载到congifmap
- 使用secret 编码存储数据库敏感信息 不会暴漏password等
- 使用动态存储自动创建 PV 默认删除会清掉数据
- service 有俩种：内部和对外暴漏访问连接使用
- 资源互相关联，会自动一起创建，删除其中一个其他关联资源都会删除
- 资源发生更改会发送通知信息

### 准备工作

1、 准备StorangeClass，在config目录->initconfig目录下准备好了
yaml文件按提示修改部分内容创建storageclass实现动态存储，

2、根据配置在 config->samples下有yaml文件按照提示配置参数

3、参考连接：https://zhuanlan.zhihu.com/p/375333739

4、 拉取代码给bin下面俩个权限
```shell
git clone https://github.com/LiuXiangBiao/mysql-operator.git
```
```shell
cd mysql
```
```shell
chmod 755 bin/*
```


# mysql
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/mysql:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/mysql:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023 liuxiangbiao.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

