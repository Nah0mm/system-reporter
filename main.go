package main

import (
	"fmt"
	"net/http"
	"system-reporter/getstatus"
)

func main() {

	http.HandleFunc("/status", getstatus.GetStatusHandler)
	//Start the server on port 9090
	fmt.Println("Starting the server on port :9090")
	http.ListenAndServe(":9090", nil)
}
