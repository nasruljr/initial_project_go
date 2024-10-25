package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type RequestResponseRecords struct {
	Records []any `json:"records"`
}

const max_retry = 3

func Post(ctx context.Context, destUrl *string, requestType string, headers *map[string]any, payload *map[string]any, counter int) (any, error) {
	availType := []string{"json", "form_params"}
	if !InArray(requestType, availType) {
		return nil, errors.New("invalid request type")
	}

	if counter > max_retry {
		return nil, errors.New("reach retry limit")
	}

	var err error
	var request *http.Request
	if strings.TrimSpace(requestType) == "json" {
		var jsonData []byte
		jsonData, err = json.Marshal(*payload)
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest("POST", *destUrl, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, err
		}
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Connection", "close")
	} else {
		var param = url.Values{}
		for key, val := range *payload {
			param.Set(key, fmt.Sprintf("%v", val))
		}
		request, err = http.NewRequest("POST", *destUrl, bytes.NewBufferString(param.Encode()))
		if err != nil {
			return nil, err
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	for key, val := range *headers {
		request.Header.Add(key, fmt.Sprintf("%v", val))
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		return Post(ctx, destUrl, requestType, headers, payload, counter+1)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var requestResponse RequestResponse
	err = json.Unmarshal(body, &requestResponse)
	if err != nil {
		return nil, err
	}

	return requestResponse, nil
}

func Get(ctx context.Context, destUrl *string, counter int) (any, error) {
	if counter > max_retry {
		return nil, errors.New("reach retry limit")
	}

	var err error
	var request *http.Request
	request, err = http.NewRequest("GET", *destUrl, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		return Get(ctx, destUrl, counter+1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return result, nil
}
