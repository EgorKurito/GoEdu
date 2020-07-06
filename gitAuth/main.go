package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	jsonData := map[string]string{
		"query": `
				{
					viewer {
						login
					}
				}
			`,
	}
	jsonValue, _ := json.Marshal(jsonData)

	url := "https://api.github.com/graphql"
	method := "POST"

	payload := bytes.NewBuffer(jsonValue)

	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	req.Header.Add("Authorization", "Bearer 52950be51c8027fd9a976f9496cde710981ff4f8")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
