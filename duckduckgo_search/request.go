package duckduckgo_search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

func Request(method string, requestURL string, params map[string]interface{}, data map[string]interface{}) (*[]byte, error) {
	versions := []string{"Chrome/120.0.0.0", "Chrome/121.0.0.0", "Chrome/122.0.0.0",
		"Chrome/123.0.0.0", "Chrome/124.0.0.0", "Chrome/125.0.0.0",
		"Chrome/126.0.0.0", "Chrome/127.0.0.0", "Chrome/128.0.0.0",
		"Chrome/129.0.0.0"}
	version := versions[rand.Intn(len(versions))]

	var request *http.Request
	var err error

	if method == "GET" {

		// setting query param
		queryParams := url.Values{}
		for key, value := range params {
			if value == "" {
				continue
			}
			queryParams.Add(key, value.(string))
		}
		requestURL = fmt.Sprintf("%s?%s", requestURL, queryParams.Encode())

		// set new request
		request, err = http.NewRequest(method, requestURL, nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to create new request")
		}

	} else if method == "POST" {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal mody data")
		}
		bodyData := bytes.NewBuffer(jsonData)
		request, err = http.NewRequest(method, requestURL, bodyData)
		if err != nil {
			return nil, fmt.Errorf("Failed to create new request")
		}
	} else {
		return nil, fmt.Errorf("Wrong Request")
	}

	// set header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "text/html")
	request.Header.Set("User-Agent", fmt.Sprintf("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) %s Safari/537.36", version))

	// send request
	var client = http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	// Read response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	// Check status code
	if response.StatusCode != 200 {
		return nil, errors.New("Status code: " + strconv.Itoa(response.StatusCode) + " Body: " + string(bodyBytes))
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	return &bodyBytes, nil
}
