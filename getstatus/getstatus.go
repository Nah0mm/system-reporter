package getstatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"system-reporter/converter"
	"system-reporter/notification"
	"time"
)

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
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
	var statusList []string
	for info := range results {
		statusList = append(statusList, info)
		notification.ExecuteSend(ereceiver)
		notification.ExecuteSend(smsreceiver)
	}
	//Setting header to json
	w.Header().Set("Content-Type", "application/json")
	//Encoding the slice to json
	json.NewEncoder(w).Encode(statusList)

	fmt.Println("Scan finished in ", time.Since(startTime))
}
