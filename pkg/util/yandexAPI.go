package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var uri string = "https://speller.yandex.net/services/spellservice.json/checkText?text="

type SpellingError struct {
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

type ErrorResponse struct {
	Errors []SpellingError `json:"errors"`
}

func SpellerApi(text string) ([]SpellingError, error) {
	text = strings.ReplaceAll(text, " ", "+")
	newUri := uri + text

	req, err := http.NewRequest("GET", newUri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	apiBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from spelling API: %w", err)
	}

	var spellingErrors []SpellingError
	if err := json.Unmarshal(apiBody, &spellingErrors); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return spellingErrors, nil
}
