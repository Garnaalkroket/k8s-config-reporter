package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func initConfig() *rest.Config {

	authConfig := &rest.Config{}

	AUTH_METHOD := os.Getenv("AUTH_METHOD")

	if AUTH_METHOD == "kubeconfig" {
		fmt.Printf("Using kubeconfig authentication. \n")
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		authConfig = config
	}

	if AUTH_METHOD == "cluster" {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		authConfig = config
	}
	return authConfig
}
