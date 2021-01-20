package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddDocument(functionName string, data map[string]string) {
	URL := "db.storagesmash.com"
	postBody, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	buffedBody := bytes.NewBuffer(postBody)

	response, err := http.Post(URL+functionName, "application/json", buffedBody)

	fmt.Printf("[RESPONSE] %s", response.Body)

	// if response.Request.Body.Read()
}
