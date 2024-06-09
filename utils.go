package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

type Figma struct {
	FILE_KEY     string
	COMPONENT_ID string
	API_KEY      string
}

func (figma *Figma) getUri() (string, error) {
	component_url := `https://api.figma.com/v1/files/{{.FILE_KEY}}/nodes?ids={{.COMPONENT_ID}}`
	t, parsingFailure := template.New("figma_component_uri").Parse(component_url)
	if parsingFailure != nil {
		return "", parsingFailure
	}

	var result bytes.Buffer

	err := t.Execute(&result, figma)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func (figma *Figma) FetchComponent() (Component, error) {
	var resp Component

	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	uri, uriError := figma.getUri()
	if uriError != nil {
		return resp, uriError
	}

	req, requestError := http.NewRequest("GET", uri, nil)
	if requestError != nil {
		return resp, requestError
	}

	// Set the custom header
	req.Header.Set("X-Figma-Token", figma.API_KEY)

	httpResp, httpError := client.Do(req)
	if httpError != nil {
		return resp, httpError
	}
	// Ensure the response body is closed after we're done with it
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("HTTP status code is %w", httpResp.StatusCode)
	}

	body, readBodyError := ioutil.ReadAll(httpResp.Body)
	if readBodyError != nil {
		return resp, readBodyError
	}

	if unmarshallingError := json.Unmarshal(body, &resp); unmarshallingError != nil {
		return resp, unmarshallingError
	}

	return resp, nil
}
