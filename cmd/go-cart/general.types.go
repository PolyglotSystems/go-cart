package goCart

import (
	"time"
)

/*====================================================================================================
  Application Types
====================================================================================================*/

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// Config struct for webapp config
type Config struct {
	GoCart ConfigYAML `yaml:"goCart"`
}

// ConfigYAML is the overall specification for what is defined for this go-cart server
type ConfigYAML struct {
	// CachePath is the path to the cache directory where generated metadata are stored
	CachePath string `yaml:"cache_path"`

	// KubernetesAuthentication is the authentication configuration for the kubernetes API
	KubernetesAuthentication KubernetesAuthentication `yaml:"k8s_auth"`

	// Server is the server configuration structure
	Server  Server `yaml:"server"`
}

// KubernetesAuthentication is the configuration for the Kubernetes API Authentication
type KubernetesAuthentication struct {
	// Method is the authentication method to use
	// Valid values are: "kubeconfig", "serviceaccount"
	Method string `yaml:"method"`

	// Paths is the set of paths that could be set depending on the method
	Paths KubernetesAuthenticationPaths `yaml:"paths"`
}

// KubernetesAuthenticationPaths is the set of paths that could be set depending on the method
type KubernetesAuthenticationPaths struct {
	// KubeConfigPath is the path to the kubeconfig file
	KubeConfigPath string `yaml:"kubeconfig_path,omitempty"`
	// ServiceAccountPath is the path to the service account file
	ServiceAccountPath string `yaml:"serviceaccount_path,omitempty"`
	// CertificatePath is the path to the certificate file
	CertificatePath string `yaml:"certificate_path,omitempty"`
}

// Server configures the HTTP server
type Server struct {
	// Host is the local machine IP Address to bind the HTTP Server to (eg: 0.0.0.0)
	Host string `yaml:"host"`

	// Port is the local machine TCP Port to bind the HTTP Server to
	Port    string `yaml:"port"`

	// BasePath is the API base path (eg: /api)
	BasePath string `yaml:"base_path"`

	// Timeout is the timeout for the HTTP server
	Timeout struct {
		// Server is the general server timeout to use for graceful shutdowns
		Server time.Duration `yaml:"server"`

		// Write is the amount of time to wait until an HTTP server write opperation is cancelled
		Write time.Duration `yaml:"write"`

		// Read is the amount of time to wait until an HTTP server read operation is cancelled
		Read time.Duration `yaml:"read"`

		// Idle is the amount of time to wait until an IDLE HTTP session is closed
		Idle time.Duration `yaml:"idle"`
	} `yaml:"timeout"`
}

// ReturnGenericMessage - Generic message structure containing a status ID, error messages, and return messages
type ReturnGenericMessage struct {
	Status   string   `json:"status"`
	Errors   []string `json:"errors"`
	Messages []string `json:"messages"`
}