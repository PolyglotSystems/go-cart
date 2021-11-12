package goCart

import (

	"k8s.io/client-go/kubernetes"
)

var goCartVersion string = "0.0.1"
var readConfig *Config

var kClient *kubernetes.Clientset

// BUFFERSIZE is for copying files
var BUFFERSIZE int64 = 4096 // 4096 bits = default page size on OSX

const serverUA = "GoCart/0.0.1"