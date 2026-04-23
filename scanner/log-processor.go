package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Device struct {
	Name string `json:"name"`
	IP   string `json:"IP"`
}

func (d *Device) PingDevice(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	timeout := time.Duration(2 * time.Second)
	conn, err := net.DialTimeout("tcp", d.IP+":9001", timeout)
	if err != nil {
		ch <- fmt.Sprintf("Device %s not responding %v", d.IP, err.Error())
		return
	}
	ch <- fmt.Sprintf("Device %s is up", d.IP)
	conn.Close()

}
