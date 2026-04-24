package converter

import (
	"encoding/json"
	"flag"
	"os"
	"system-reporter/scanner"
)

func Converter() ([]scanner.Device, error) {
	configPath := flag.String("config", "devices.json", "Path to the devices configuration file")
	flag.Parse()
	data, err := os.ReadFile(*configPath)
	if err != nil {

		return nil, err
	}
	var devices []scanner.Device
	err = json.Unmarshal(data, &devices)
	if err != nil {

		return nil, err
	}

	return devices, nil
}
