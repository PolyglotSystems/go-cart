package goCart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/klog/v2"
)

// getDashboardData returns the application availability & health
func getDashboardData(w http.ResponseWriter, r *http.Request) {
	logNeworkRequestStdOut(r.Method+" "+r.RequestURI, r)

	namespaces, err := QueryNamespaces(kClient, "")
	if err != nil {
		klog.Errorf("Error querying namespaces: %v", err)
	}
	numberOfNamespaces := len(namespaces.Items)

	//pods, err := c.QueryPods(kClient, "default", "app=go-cart")
	pods, err := QueryPods(kClient, "", "")
	if err != nil {
		klog.Errorf("Error querying pods: %v", err)
	}
	numberOfPods := len(pods.Items)

	services, err := QueryServices(kClient, "", "")
	if err != nil {
		klog.Errorf("Error querying services: %v", err)
	}
	numberOfServices := len(services.Items)

	data := make(map[string]string)
	data["numberOfNamespaces"] = fmt.Sprintf("%d", numberOfNamespaces)
	data["numberOfPods"] = fmt.Sprintf("%d", numberOfPods)
	data["numberOfServices"] = fmt.Sprintf("%d", numberOfServices)

	returnData := RESTGETDasboardDataJSONReturn{
		Status:   "OK",
		Errors:   []string{},
		Messages: []string{},
		Data:     data,
	}
	returnResponse, _ := json.Marshal(returnData)
	fmt.Fprintf(w, string(returnResponse))
}

func getNamespacesQuery(w http.ResponseWriter, r *http.Request) {
	logNeworkRequestStdOut(r.Method+" "+r.RequestURI, r)

	namespaces, err := QueryNamespaces(kClient, "")
	if err != nil {
		klog.Errorf("Error querying namespaces: %v", err)
	}

	data := []string{}
	for _, namespace := range namespaces.Items {
		data = append(data, namespace.Name)
	}

	returnData := RESTGETNamespaceJSONReturn{
		Status:   "OK",
		Errors:   []string{},
		Messages: []string{},
		Data:     data,
	}
	returnResponse, _ := json.Marshal(returnData)
	fmt.Fprintf(w, string(returnResponse))

}
