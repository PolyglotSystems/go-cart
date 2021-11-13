package gocart

import (
	"encoding/json"

	"k8s.io/klog/v2"
)

func runMapper() []byte {
	// Get the namespaces
	namespaces, err := QueryNamespaces(kClient, "")
	if err != nil {
		klog.Errorf("Error querying namespaces: %s", err)
		return []byte{}
	}
	//var nsMap map[string][]ServiceData
	nsMap := make(map[string][]ServiceData)

	// Loop through the namespaces and get the services
	for _, namespace := range namespaces.Items {
		// Get the services
		services, err := QueryServices(kClient, namespace.Name, "")
		if err != nil {
			klog.Errorf("Error querying services in namespace %s: %s", namespace.Name, err)
			continue
		}
		svcCollection := []ServiceData{}
		for _, services := range services.Items {
			svcCollection = append(svcCollection, ServiceData{
				ServiceName:      services.Name,
				Namespace:        services.Namespace,
				Labels:           services.Labels,
				Annotations:      services.Annotations,
				ServiceSelectors: services.Spec.Selector,
			})
		}
		// With all the services discovered, add them to the map of namespaces by this namespace ID
		nsMap[namespace.Name] = svcCollection
	}

	// Convert the map to JSON
	json, err := json.Marshal(nsMap)
	if err != nil {
		klog.Errorf("Error marshalling map to JSON: %s", err)
		return []byte{}
	}

	return json
}
