package converter

import (
	"encoding/json"
	"os"
	"system-reporter/scanner"
)

func Converter() ([]scanner.Device, error) {
	data, err := os.ReadFile("devices/devices.json")
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
