package config

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Kubernetes struct {
	k8sClient *kubernetes.Clientset
}

func (c *Kubernetes) Client() *kubernetes.Clientset {
	return c.k8sClient
}

func LoadKubernetes() (*Kubernetes, error) {
	inClusterConf, err := rest.InClusterConfig()
	if err != nil {
		log.Print("failed when get in-cluster config")
		return nil, err
	}

	k8sClient, err := kubernetes.NewForConfig(inClusterConf)
	if err != nil {
		return nil, err
	}

	conf := Kubernetes{
		k8sClient: k8sClient,
	}

	return &conf, nil
}

func LoadKubectlKubernetes() (*Kubernetes, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	conf := Kubernetes{
		k8sClient: k8sClient,
	}

	return &conf, nil
}
