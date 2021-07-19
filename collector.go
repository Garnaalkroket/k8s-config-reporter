package main

import (
	"fmt"

	"k8s.io/client-go/rest"

	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type resourceCollector struct {
	resourceMetric *prometheus.Desc
	kubeConfig     *rest.Config
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newResourceCollector(config *rest.Config) *resourceCollector {
	return &resourceCollector{
		resourceMetric: prometheus.NewDesc("k8s_profile_metrics",
			"Tracks different resource profiles through labels",
			[]string{"namespace", "profile", "status"}, nil,
		),
		kubeConfig: config,
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *resourceCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.resourceMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *resourceCollector) Collect(ch chan<- prometheus.Metric) {
	profileMap := resources()
	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	//ch <- prometheus.MustNewConstMetric(collector.gitMetric, prometheus.GaugeValue, my_time.time, my_time.name)
	for k1, v1 := range profileMap {
		for k2, v2 := range v1 {
			fmt.Printf("Creating metric for %s -> %s -> %s \n", k1, k2, v2)
			ch <- prometheus.MustNewConstMetric(collector.resourceMetric, prometheus.GaugeValue, 1, k1, k2, v2)
		}
	}
}
