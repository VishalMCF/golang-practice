package main

import (
	"fmt"
	"net/http"
)

type KeyValueServer struct {
	d *DataStore
}

func (server KeyValueServer) get(w http.ResponseWriter, req *http.Request) {

}

func (server KeyValueServer) multiply(w http.ResponseWriter, req *http.Request) {

}

func (server KeyValueServer) increment(w http.ResponseWriter, req *http.Request) {

}

func (server KeyValueServer) add(w http.ResponseWriter, req *http.Request) {

}

func main() {
	fmt.Println("Hello key value server. Thanks to demonstrate concurrency in golang")
}
