package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var config *rest.Config = &rest.Config{}

var resourceMap map[string]map[string]string

func main() {

	//Allocate memory
	resourceMap = make(map[string]map[string]string)

	//Set authentication configuration
	config = initConfig()

	//Initialize
	conf()

	//Refresh using goroutine
	go queryConfig(conf())

	//Start Metrics and plain JSON endpoint
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/resources", resourceHandler)
	http.HandleFunc("/", indexHandler)
	resourceMetrics := newResourceCollector(conf())
	prometheus.MustRegister(resourceMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func conf() *rest.Config {
	return config
}

func resources() map[string]map[string]string {
	return resourceMap
}
