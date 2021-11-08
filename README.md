# kube-operator
###  安装operator-sdk 脚手架工具
教程参考： https://sdk.operatorframework.io/docs/installation/

### 使用教程
- 1. 初始话项目
```sh
operator-sdk init --domain mdemo.com --repo github.com/humorliang/kube-operator
```
- 1.1 修改project info
```
go.mod  中的go 版本修改为 开发环境的go 版本，建议>=15

makefile 文件中的镜像仓库管理位置 ：
IMAGE_TAG_BASE ?= 192.168.15.128:5000/kube-operator
```
- 更新go版本
```sh
# How to update the Go version

#System: Debian/Ubuntu/Fedora. Might work for others as well.
#1. Uninstall the exisiting version

# As mentioned here, to update a go version you will first need to uninstall the original version.

# To uninstall, delete the /usr/local/go directory by:

$ sudo rm -rf /usr/local/go

# 2. Install the new version
# addr:  https://golang.org/dl/

# Go to the downloads page and download the binary release suitable for your system.
# 3. Extract the archive file

# To extract the archive file:

$ sudo tar -C /usr/local -xzf go1.16.9.linux-amd64.tar.gz

# 4. Make sure that your PATH contains /usr/local/go/bin

$ echo $PATH | grep "/usr/local/go/bin"
```
特定集群的依赖mod 配置
```
require (
	github.com/operator-framework/operator-sdk v0.15.2
	sigs.k8s.io/controller-runtime v0.4.0
)
// Pinned to kubernetes-1.16.2
replace (
	k8s.io/api => k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191016113550-5357c4baaf65
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191016112112-5190913f932d
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191016114015-74ad18325ed5
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191016115326-20453efc2458
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115129-c07a134afb42
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191004115455-8e001e5d1894
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191016111319-039242c015a9
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191016115521-756ffa5af0bd
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112429-9587704a8ad4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114939-2b2b218dc1df
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114407-2e83b6f20229
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114748-65049c67a58b
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191016120415-2ed914427d51
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114556-7841ed97f1b2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115753-cf0698c3a16b
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191016113814-3b1a734dba6e
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112829-06bb3c9d77c9
)
replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309 // Required by Helm
replace github.com/openshift/api => github.com/openshift/api v0.0.0-20190924102528-32369d4db2ad // Required until https://github.com/operator-framework/operator-lifecycle-manager/pull/1241 is resolved
```

- 2. 创建自定义的API 

```sh
# 指定group version 这个资源的kind  并且 生成这个资源 和 controller
operator-sdk create api --group mapp --version v1alpha1 --kind Mdemo --resource --controller
```
目录结构
```sh
├── api   # CRD 结构对象目录
│   └── v1alpha1
│       ├── groupversion_info.go
│       ├── mdemo_types.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── controller-gen
├── config  # 相关资源的 yaml 文件
│   ├── crd
│   │   ├── kustomization.yaml
│   │   ├── kustomizeconfig.yaml
│   │   └── patches
│   │       ├── cainjection_in_mdemoes.yaml
│   │       └── webhook_in_mdemoes.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_auth_proxy_patch.yaml
│   │   └── manager_config_patch.yaml
│   ├── manager
│   │   ├── controller_manager_config.yaml
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── manifests
│   │   └── kustomization.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── auth_proxy_client_clusterrole.yaml
│   │   ├── auth_proxy_role_binding.yaml
│   │   ├── auth_proxy_role.yaml
│   │   ├── auth_proxy_service.yaml
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── mdemo_editor_role.yaml
│   │   ├── mdemo_viewer_role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   ├── samples
│   │   ├── kustomization.yaml
│   │   └── mapp_v1alpha1_mdemo.yaml
│   └── scorecard
│       ├── bases
│       │   └── config.yaml
│       ├── kustomization.yaml
│       └── patches
│           ├── basic.config.yaml
│           └── olm.config.yaml
├── controllers  # controller 控制器的代码目录
│   ├── mdemo_controller.go
│   └── suite_test.go
├── Dockerfile
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── main.go  # 入口文件
├── Makefile
├── PROJECT
└── README.md
```

- 3. 编写逻辑
