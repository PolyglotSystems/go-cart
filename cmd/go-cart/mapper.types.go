package goCart

type NamespacedServiceData struct {
	Namespace string
	Services  []ServiceData
}

type ServiceData struct {
	ServiceName      string
	Namespace        string
	ServiceSelectors map[string]string
	Labels           map[string]string
	Annotations      map[string]string
}
