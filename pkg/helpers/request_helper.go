package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ResolveUrl(resource string) string {
	return "https://jsonplaceholder.typicode.com" + resource
}

func SendRequest(method, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ReadResponseBody(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func HandlePostResponse(response *http.Response, target interface{}) (int, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}
	defer response.Body.Close()

	if err := json.Unmarshal(body, target); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body to map: %w", err)
	}

	if id, ok := responseBody["id"].(float64); ok {
		return int(id), nil
	}

	return 0, fmt.Errorf("response does not contain an id property")
}
