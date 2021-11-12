package goCart

import (
	"os"
	"io/ioutil"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"net"
	certutil "k8s.io/client-go/util/cert"
)

// authenticateViaKubeConfig authenticates via kubeconfig
// @param kubeconfigPath string
// @return (*kubernetes.Clientset, error)
func authenticateViaKubeConfig(kubeConfigPath string) (*kubernetes.Clientset, error) {
	
	// Create a Config (k8s.io/client-go/rest.Config)
	// from the kubeConfigPath
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	// Create a new Clientset which will use the Config
	// to communicate with the apiserver.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
	
// authenticateViaServiceAccount returns a Clientset object which uses the service account
// kubernetes gives to pods. It's intended for clients that expect to be
// running inside a pod running on kubernetes. 
// @param serviceAccountTokenPath string, path to the service account token
// @param serviceAccountCertificatePath string, path to the service account certificate
// @return (*kubernetes.Clientset, error)
func authenticateViaServiceAccount(serviceAccountTokenPath string, serviceAccountCertificatePath string) (*kubernetes.Clientset, error) {

	// const (
	// 	tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	// 	rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	// )

	host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
	if len(host) == 0 || len(port) == 0 {
		return nil, rest.ErrNotInCluster
	}

	token, err := ioutil.ReadFile(serviceAccountTokenPath)
	if err != nil {
		return nil, err
	}

	tlsClientConfig := rest.TLSClientConfig{}

	if _, err := certutil.NewPool(serviceAccountCertificatePath); err != nil {
		klog.Errorf("Expected to load root CA config from %s, but got err: %v", serviceAccountCertificatePath, err)
	} else {
		tlsClientConfig.CAFile = serviceAccountCertificatePath
	}

	// Create the API Configuration structure
	config := &rest.Config{
		// TODO: switch to using cluster DNS.
		Host:            "https://" + net.JoinHostPort(host, port),
		TLSClientConfig: tlsClientConfig,
		BearerToken:     string(token),
		BearerTokenFile: serviceAccountTokenPath,
	}

	// Create a new Clientset which will use the Config
	// to communicate with the apiserver.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}