package goCart

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"

	//"strings"
	"syscall"
	"time"

	"k8s.io/klog/v2"
)

// NewRouter generates the router used in the HTTP Server
func NewRouter(basePath string) *http.ServeMux {

	var formattedBasePath string
	var apiPrefix string

	if basePath == "" {
		basePath = "/go-cart"
	}
	formattedBasePath = strings.TrimRight(basePath, "/")

	// Create router and define routes and return that router
	router := http.NewServeMux()

	//====================================================================================
	// TEST ENDPOINT
	// Test out a random function maybe
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		returnData := &ReturnGenericMessage{
			Status:   "test",
			Errors:   []string{},
			Messages: []string{"Test!"}}
		returnResponse, _ := json.Marshal(returnData)
		fmt.Fprintf(w, string(returnResponse))
	})

	//====================================================================================
	// KUBERNETES ENDPOINTS
	// Version Output - reads from variables.go
	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		applicationVersionAPI(w, r)
	})

	// Healthz endpoint for kubernetes platforms
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthZAPI(w, r)
	})

	//====================================================================================
	// START V1 API
	//====================================================================================
	apiPrefix = "/v1"

	router.HandleFunc(formattedBasePath+apiPrefix+"/namespaces", func(w http.ResponseWriter, r *http.Request) {
		logNeworkRequestStdOut(r.Method+" "+r.RequestURI, r)
		switch r.Method {
		case "GET":
			// index - get namespaces
			getNamespacesQuery(w, r)
		default:
			methodNotAllowedAPI(w, r)
		}
	})

	router.HandleFunc(formattedBasePath+apiPrefix+"/dashboard", func(w http.ResponseWriter, r *http.Request) {
		logNeworkRequestStdOut(r.Method+" "+r.RequestURI, r)
		switch r.Method {
		case "GET":
			// index - get dashboard data
			getDashboardData(w, r)
		default:
			methodNotAllowedAPI(w, r)
		}
	})

	return router
}

// RunHTTPServer will run the HTTP Server
func (config Config) RunHTTPServer() {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.GoCart.Server.Timeout.Server,
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         config.GoCart.Server.Host + ":" + config.GoCart.Server.Port,
		Handler:      NewRouter(config.GoCart.Server.BasePath),
		ReadTimeout:  config.GoCart.Server.Timeout.Read * time.Second,
		WriteTimeout: config.GoCart.Server.Timeout.Write * time.Second,
		IdleTimeout:  config.GoCart.Server.Timeout.Idle * time.Second,
	}

	// Only listen on IPV4
	l, err := net.Listen("tcp4", config.GoCart.Server.Host+":"+config.GoCart.Server.Port)
	check(err)

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Alert the user that the server is starting
	klog.Infof("Server is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		//if err := server.ListenAndServe(); err != nil {
		if err := server.Serve(l); err != nil {
			if err == http.ErrServerClosed {
				// Normal interrupt operation, ignore
			} else {
				klog.Fatalf("Server failed to start due to err: %v", err)
			}
		}

	}()

	//go func() {
	//	//testQueryLoop()
	//	for {
	//		klog.Infof(`{"namespaces": %v }`, string(runMapper()))
	//		time.Sleep(time.Second * 10)
	//	}
	//}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	klog.Infof("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		klog.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
