package gocart

// NamespacedServiceData provides the top-down view of the service data starting at the Namespace
type NamespacedServiceData struct {
	Namespace string
	Services  []ServiceData
}

// ServiceData provides key information about a service
type ServiceData struct {
	ServiceName      string
	Namespace        string
	ServiceSelectors map[string]string
	Labels           map[string]string
	Annotations      map[string]string
}
