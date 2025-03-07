package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchNFTMetaData performs a GET request to fetch JSON data from the provided URL
func FetchNFTMetaData(url string) (map[string]any, error) {
	// Создаем HTTP GET запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Распарсим JSON данные в map
	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return data, nil
}
