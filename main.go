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
	phone := notification.SMS{PhoneNumber: "25187459687"}
	email := notification.Email{Address: "someone@example.com"}
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
		fmt.Printf("Sent %s to %s\n", info, phone.PhoneNumber)
		fmt.Printf("Sent %s to %s\n", info, email.Address)
	}

	fmt.Println("Scan finished in ", time.Since(startTime))
}
