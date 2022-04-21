package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func sendRequest12() {

	json := []byte(`{"id": "1234","name": "ya"}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("POST", "http://192.168.31.197:8000/delay", body)

	// Headers
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
		return
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func main() {
	sendRequest12()
}
