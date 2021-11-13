package gocart

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
)

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	checkAndFail(err)
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	readConfig = config

	return config, nil
}

// PreflightSetup just makes sure the stage is set
func PreflightSetup() {
	// Set up Kubernetes Authentication
	switch readConfig.GoCart.KubernetesAuthentication.Method {
	case "kubeconfig":
		// Set up Kubernetes Authentication via KubeConfig
		klog.Infof("Connecting with kubeconfig located at %s", readConfig.GoCart.KubernetesAuthentication.Paths.KubeConfigPath)
		newClient, err := authenticateViaKubeConfig(readConfig.GoCart.KubernetesAuthentication.Paths.KubeConfigPath)
		if err != nil {
			klog.Errorf("Error authenticating via kubeconfig:", err)
			os.Exit(1)
		}
		kClient = newClient
	case "serviceaccount":
		// Set up Kubernetes Authentication via ServiceAccount
		klog.Infof("Connecting with ServiceAccount located at %s", readConfig.GoCart.KubernetesAuthentication.Paths.ServiceAccountPath)
		newClient, err := authenticateViaServiceAccount(readConfig.GoCart.KubernetesAuthentication.Paths.ServiceAccountPath, readConfig.GoCart.KubernetesAuthentication.Paths.CertificatePath)
		if err != nil {
			klog.Errorf("Error authenticating via ServiceAccount:", err)
			os.Exit(1)
		}
		kClient = newClient
	default:
		klog.Errorf("Kubernetes Authentication Method '%s' is not supported", readConfig.GoCart.KubernetesAuthentication.Method)
		os.Exit(1)
	}
	klog.Infof("%s", "Preflight complete!")
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// rmStrFromStrSlice removes a string from a string slice
func rmStrFromStrSlice(r string, s []string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// B64EncodeBytesToStr converts a byte slice to a Base64 Encoded String
func B64EncodeBytesToStr(input []byte) string {
	return b64.StdEncoding.EncodeToString(input)
}

// B64DecodeBytesToBytes converts a Base64 byte slice to a Base64 Decoded Byte slice
func B64DecodeBytesToBytes(input []byte) ([]byte, error) {
	return B64DecodeStrToBytes(string(input))
}

// B64DecodeStrToBytes converts a Base64 string to a Base64 Decoded Byte slice
func B64DecodeStrToBytes(input string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(input)
}
