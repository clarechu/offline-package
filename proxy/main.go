package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	port             = flag.String("port", ":8080", "The port to listen on")
	proxyAddress     = flag.String("proxy_address", ":8081", "The address of the proxy server")
	defaultTransport *http.Transport
)

func init() {
	defaultTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
func main() {
	flag.Parse()
	http.HandleFunc("/", proxy)

	err := http.ListenAndServe(*port, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Listening on " + *port)
}

func proxy(w http.ResponseWriter, r *http.Request) {
	r.URL.Scheme = "http"
	r.URL.Host = *proxyAddress
	trip, err := defaultTransport.RoundTrip(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, fmt.Errorf(`{"message": "%s"}`, err.Error()))
		return
	}
	if trip.StatusCode != 200 {
		data, err := io.ReadAll(trip.Body)
		if err != nil {
			WriteError(w, http.StatusBadRequest, fmt.Errorf(`{"message": "%s"}`, err.Error()))
			return
		}
		WriteError(w, trip.StatusCode, errors.New(string(data)))
		return
	}
	data, err := io.ReadAll(trip.Body)
	if err != nil {
		WriteError(w, http.StatusBadRequest, fmt.Errorf(`{"message": "%s"}`, err.Error()))
		return
	}
	for k, vs := range trip.Header {
		for _, v := range vs {
			w.Header().Set(k, v)
		}
	}
	w.Write(data)
}

func WriteError(w http.ResponseWriter, httpCode int, err error) {
	w.WriteHeader(httpCode)
	w.Write([]byte(err.Error()))
}
