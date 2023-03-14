package kube

import (
	"gitee.com/plutoccc/devops_app/internal/core/settings"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientset(cluster string) (client kubernetes.Interface, cfg *rest.Config, err error) {

	pm := settings.NewSettingManager()
	resp, err := pm.GetIntegrateSettingByName(cluster, settings.KubernetesType)
	if err != nil {
		return nil, nil, err
	}
	return buildK8sClient(resp.IntegrateSettingReq.Config.(*settings.KubeConfig))
}

func buildK8sClient(kube *settings.KubeConfig) (client kubernetes.Interface, cfg *rest.Config, err error) {
	var k8sConfig *rest.Config
	switch kube.Type {
	case settings.KubernetesConfig:
		k8sConfig, err = clientcmd.RESTConfigFromKubeConfig([]byte(kube.Conf))
		if err != nil {
			return nil, nil, err
		}
	case settings.KubernetesToken:
		k8sConfig = &rest.Config{
			BearerToken:     kube.Conf,
			TLSClientConfig: rest.TLSClientConfig{Insecure: true},
			Host:            kube.URL,
		}
	}

	clientSet, err := kubernetes.NewForConfig(k8sConfig)
	return clientSet, k8sConfig, err
}
