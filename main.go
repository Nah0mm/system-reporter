package main

import (
	"fmt"
	"sync"
	"system-reporter/converter"
	"time"
)

func main() {
	startTime := time.Now()
	devices, err := converter.Converter()
	if err != nil {
		fmt.Println("Error occcured", err)
		return
	}
	var wg sync.WaitGroup
	results := make(chan string)
	for _, device := range devices {
		wg.Add(1)
		go (&device).PingDevice(results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for info := range results {
		fmt.Println(info)
	}
	fmt.Println("Scan finished in ", time.Since(startTime))
}
