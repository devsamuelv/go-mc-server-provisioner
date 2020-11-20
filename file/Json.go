package file

import (
	"fmt"
	"os"
	"io/ioutil"
	"main/types"
	"encoding/json"
)

// ReadConfigFile Reads the Server Config File
func ReadConfigFile(file string) ([]string) {
	var data types.DriveConfig;
	jsonFile, err := os.Open(file);

	if err != nil {
		fmt.Println(err);
	}

	defer jsonFile.Close();

	byteValue, _ := ioutil.ReadAll(jsonFile);

	json.Unmarshal(byteValue, &data);

	return data.Drives;
}