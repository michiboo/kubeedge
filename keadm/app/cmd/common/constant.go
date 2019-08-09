/*
Copyright 2019 The Kubeedge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

const (
	//KubeEdgeVersion sets the version of KubeEdge to be used
	KubeEdgeVersion = "kubeedge-version"

	// K8SImageRepository sets the image repository of Kubernetes
	K8SImageRepository = "image-repository"

	//K8SPodNetworkCidr sets pod network cidr of Kubernetes
	K8SPodNetworkCidr = "pod-network-cidr"

	//DockerVersion sets the version of Docker to be used
	DockerVersion = "docker-version"

	//KubernetesVersion sets the version of Kuberneted to be used
	KubernetesVersion = "kubernetes-version"

	//KubeConfig sets the path of kubeconfig , default is ""
	KubeConfig = "kube-config"

	//K8SAPIServerIPPort sets the IP:Port of Kubernetes api-server
	K8SAPIServerIPPort = "k8sserverip"

	//CloudCoreIP sets the IP of KubeEdge cloud component
	CloudCoreIP = "cloudcoreip"

	//KubeEdge Node unique idenfitcation string
	EdgeNodeID = "edgenodeid"

	//KubeEdge interface name string
	InterfaceName = "interfacename"

	//CertPath sets the path of the certificates generated by the KubeEdge Cloud component
	CertPath = "certPath"

	//DefaultCertPath is the default certificate path in edge node
	DefaultCertPath = "/etc/kubeedge/certs"

	//DefaultKubeEdgeVersion is the current default version of KubeEdge
	DefaultKubeEdgeVersion = "0.3.0"

	//DefaultDockerVersion is the current default version of Docker
	DefaultDockerVersion = "18.06.0"

	//DefaultK8SVersion is the current default version of K8S
	DefaultK8SVersion = "1.14.1"

	// DefaultProjectID is default project id
	DefaultProjectID = "e632aba927ea4ac2b575ec1603d56f10"

	//DefaultImageRepository is the default k8s image repository
	DefaultK8SImageRepository = "k8s.gcr.io"

	//DefaultPodNetworkCidr is the default k8s pod network cidr
	DefaultK8SPodNetworkCidr = "100.64.0.0/10"

	VendorK8sPrefix = "v1.10.9-kubeedge-"

	// RuntimeType is default runtime type
	RuntimeType = "runtimetype"
)
