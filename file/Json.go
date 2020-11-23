package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/types"
	"os"
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
