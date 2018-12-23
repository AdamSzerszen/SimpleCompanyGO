package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfiguration() Configuration {
	var configuration Configuration
	//filename is the path to the json config file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return configuration
}
