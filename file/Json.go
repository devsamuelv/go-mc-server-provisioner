package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/types"
	"os"
	"strings"
)

// ReadConfigFile Reads the Server Config File
func ReadConfigFile(file string) types.Config {
	var data types.Config
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

	return data
}

// WriteConfigFile Writes To The Server Config File
func WriteConfigFile(file string, old string, new string) {
	read, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("[ERROR] %s \n", err)
	}

	NewContents := strings.Replace(string(read), old, new, -1)

	err = ioutil.WriteFile(file, []byte(NewContents), 0)

	if err != nil {
		fmt.Printf("[ERROR] %s \n", err)
	}
}
