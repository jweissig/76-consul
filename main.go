package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/consul/api"
)

// used to dump headers for debugging
func handler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// set hostname (used for demo)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error:", err)
	}
	fmt.Fprintf(w, "Hello from %v\n\n", hostname)

	// get kv + service instances
	fmt.Fprintf(w, "KV from Consul:\n")

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	rmax, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		log.Print("kv get error: ", err)
	}
	fmt.Fprintf(w, "%v %s\n", rmax.Key, rmax.Value)

	// Lookup the pair
	rttl, _, err := kv.Get("REDIS_TIMEOUT", nil)
	if err != nil {
		log.Print("kv get error: ", err)
	}
	fmt.Fprintf(w, "%v %s\n", rttl.Key, rttl.Value)

	// get services
	services, metainfo, err := client.Health().Service("redis-cache", "global", true, &api.QueryOptions{})

	if err != nil {
		fmt.Printf("error getting instances from Consul: %v", err)
	}

	fmt.Fprintf(w, "\nRedis instances from Consul (%v):\n", len(services))

	for _, service := range services {
		fmt.Fprintf(w, "%v:%v\n", service.Service.Address, service.Service.Port)
	}

	fmt.Fprintf(w, "\nRequestTime %v\n", metainfo.RequestTime)

	return

}

var client *api.Client

func main() {

	config := api.DefaultConfig()
	config.Address = "10.138.0.10:8501"

	// Get a new client
	var clientErr error
	client, clientErr = api.NewClient(config)
	if clientErr != nil {
		log.Print("client error: ", clientErr)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5005", nil))
}
