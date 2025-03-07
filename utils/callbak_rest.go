package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func CallbackREST(url string, success any) (bool, error) {
	// Serialize struct to JSON
	body, err := json.Marshal(success)
	if err != nil {
		return false, err
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return false, err // Error occurred while sending the request
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err // Error occurred while reading the response
	}

	// Check for response code 200 and body containing "OK"
	if resp.StatusCode == http.StatusOK && strings.HasPrefix(string(responseBody), "OK") {
		return true, nil
	}

	return false, nil // Return false if not 200 or body is not "OK"
}
