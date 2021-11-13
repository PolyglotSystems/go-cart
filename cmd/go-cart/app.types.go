package gocart

// RESTGETDasboardDataJSONReturn handles the data returned by the GET /dashboard endpoint for dashboard data
type RESTGETDasboardDataJSONReturn struct {
	Status   string            `json:"status"`
	Errors   []string          `json:"errors"`
	Messages []string          `json:"messages"`
	Data     map[string]string `json:"data,omitempty"`
}

// RESTGETNamespaceJSONReturn handles the data returned by the GET /namespaces endpoint for namespace listings
type RESTGETNamespaceJSONReturn struct {
	Status   string   `json:"status"`
	Errors   []string `json:"errors"`
	Messages []string `json:"messages"`
	Data     []string `json:"data"`
}
