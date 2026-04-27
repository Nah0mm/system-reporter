package main

import (
	"fmt"
	"sync"
	"system-reporter/converter"
	"system-reporter/notification"
	"time"
)

func main() {
	startTime := time.Now()
	devices, err := converter.Converter()
	ereceiver := notification.Email{Address: "someone@gmail.com"}
	smsreceiver := notification.SMS{PhoneNumber: "25187456987"}
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
		notification.ExecuteSend(ereceiver)
		notification.ExecuteSend(smsreceiver)
	}
	fmt.Println("Scan finished in ", time.Since(startTime))
}
