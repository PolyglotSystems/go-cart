package gocart

import (
	"context"
	"time"

	//appsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	//appsV1 "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

// QueryNamespaces queries the namespaces in the cluster
// @param selector the selector to query
// @return the list of namespaces and error
func QueryNamespaces(c *kubernetes.Clientset, selector string) (*v1.NamespaceList, error) {
	return c.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
}

// QueryPods queries the pods in the cluster
// @param namespace the namespace to query
// @param selector the selector to query
// @return the list of pods and error
func QueryPods(c *kubernetes.Clientset, namespace string, selector string) (*v1.PodList, error) {
	return c.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
}

// QueryServices queries the services in the cluster
// @param namespace the namespace to query
// @param selector the selector to query
// @return the list of services and error
func QueryServices(c *kubernetes.Clientset, namespace string, selector string) (*v1.ServiceList, error) {
	return c.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
}

// QueryDeployments queries the deployments in the cluster
// @param namespace the namespace to query
// @param selector the selector to query
// @return the list of deployments and error
func QueryDeployments(c *kubernetes.Clientset, namespace string, selector string) (*appsv1.DeploymentList, error) {
	return c.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
}

// testQueryLoop queries the pods in the cluster
func testQueryLoop() {

	for {
		klog.Infof("=================== START Test Query Loop")
		namespaces, err := QueryNamespaces(kClient, "")
		if err != nil {
			klog.Errorf("Error querying namespaces: %v", err)
		} else {
			klog.Infof("Found %d namespaces", len(namespaces.Items))
		}

		//deployments, err := c.QueryDeployments(kClient, "default", "app=go-cart")
		deployments, err := QueryDeployments(kClient, "", "")
		if err != nil {
			klog.Errorf("Error querying deployments: %v", err)
		} else {
			klog.Infof("Found %d deployments", len(deployments.Items))
		}

		//pods, err := c.QueryPods(kClient, "default", "app=go-cart")
		pods, err := QueryPods(kClient, "", "")
		if err != nil {
			klog.Errorf("Error querying pods: %v", err)
		} else {
			klog.Infof("Found %d pods", len(pods.Items))
		}

		services, err := QueryServices(kClient, "", "")
		if err != nil {
			klog.Errorf("Error querying services: %v", err)
		} else {
			klog.Infof("Found %d services", len(services.Items))
		}

		klog.Infof("=================== END Test Query Loop")

		time.Sleep(time.Second * 10)
	}
}
