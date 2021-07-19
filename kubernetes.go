package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func queryConfig(config *rest.Config) {

	CONFIGMAP_NAME := os.Getenv("CONFIGMAP_NAME")
	CONFIGMAP_KEYS := os.Getenv("CONFIGMAP_KEYS")

	profileMap := make(map[string]map[string]string)

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for {

		fmt.Printf("Found %d namespaces.\n", len(namespaces.Items))

		for _, ns := range namespaces.Items {

			fmt.Printf("Checking namespace: %s. \n", ns.GetName())

			maps, err := clientset.CoreV1().ConfigMaps(ns.GetName()).List(context.TODO(), metav1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name=%s", CONFIGMAP_NAME)})

			if err != nil {
				panic(err.Error())
			}

			if len(maps.Items) == 1 {
				cm := maps.Items[0]
				nsMap := make(map[string]string)
				profileMap[ns.GetName()] = nsMap
				fmt.Printf("Found ConfigMap %s in namespace %s\n", cm.GetName(), ns.GetName())
				for _, key := range strings.Split(CONFIGMAP_KEYS, ",") {
					dataMap := cm.Data
					if data, ok := dataMap[key]; ok {
						fmt.Printf("%s set to %s\n", key, data)
						profileMap[ns.GetName()][key] = data
					}
				}
			} else {
				fmt.Printf("Namespace %s does not contain ConfigMap %s\n", ns.GetName(), CONFIGMAP_NAME)
			}

		}

		for k1, v1 := range profileMap {
			for k2, v2 := range v1 {
				fmt.Printf("%s -> %s -> %s \n", k1, k2, v2)
			}
		}

		resourceMap = profileMap

		time.Sleep(5 * time.Minute)

	}
}
