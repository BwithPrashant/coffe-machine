package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadFile marshals data from passed file into the interface passed
func ReadFile(filePath string, obj interface{}) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &obj)
	return err
}

// Send true to a channel
func CloseChannel(ch chan bool) {
	ch <- true
}
