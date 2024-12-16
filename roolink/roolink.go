package roolink

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://www.roolink.io/api/v1"
)

// RooLink represents the SDK for interacting with the RooLink API.

// NewRooLink initializes a new RooLink SDK instance.
func NewSession(apiKey, protectedURL, userAgent string) *RooLink {
	return &RooLink{
		APIKey:       apiKey,
		ProtectedURL: protectedURL,
		UserAgent:    userAgent,
	}
}

// RequestLimit fetches the API request limit.
func (r *RooLink) RequestLimit(ctx context.Context) (*RequestLimit, error) {
	url := fmt.Sprintf("%s/limit?key=%s", baseURL, r.APIKey)
	res, err := r.makeRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	requests, ok := res["requests"].(float64)
	if !ok {
		return nil, errors.New("failed to parse request limit")
	}
	return &RequestLimit{Requests: int64(requests)}, nil
}

// ParseScriptData parses the script data provided.
func (r *RooLink) ParseScriptData(ctx context.Context, scriptBody []byte) (*ScriptData, error) {
	url := fmt.Sprintf("%s/parse", baseURL)
	headers := map[string]string{"Content-Type": "text/plain"}
	//return r.makeRequest(ctx, "POST", url, []byte(scriptBody), headers)
	res, err := r.makeRequest(ctx, "POST", url, scriptBody, headers)
	if err != nil {
		return nil, err
	}
	var scriptData *ScriptData
	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &scriptData)
	if err != nil {
		return nil, err
	}

	return scriptData, nil
}

// GenerateSensorData generates sensor data for validation.
func (r *RooLink) GenerateSensorData(ctx context.Context, options *SensorPayload) (*SensorResponse, error) {
	url := fmt.Sprintf("%s/sensor", baseURL)
	data, _ := json.Marshal(options)
	resp, err := r.makeRequest(ctx, "POST", url, data)
	if err != nil {
		return nil, err
	}

	sensor, ok := resp["sensor"].(string)
	if !ok {
		return nil, errors.New("failed to parse sensor data")
	}

	return &SensorResponse{Sensor: sensor}, nil
}

// GenerateSbsdBody generates the SBSD body.
func (r *RooLink) GenerateSbsdBody(ctx context.Context, options *SbsdPayload) (*SbsdBody, error) {
	url := fmt.Sprintf("%s/sbsd", baseURL)
	data, _ := json.Marshal(options)
	resp, err := r.makeRequest(ctx, "POST", url, data)
	if err != nil {
		return nil, err
	}

	body, ok := resp["body"].(string)
	if !ok {
		return nil, errors.New("failed to parse SBSD body")
	}

	return &SbsdBody{Body: body}, nil
}

// GeneratePixelData generates pixel data.
func (r *RooLink) GeneratePixelData(ctx context.Context, options *PixelPayload) (string, error) {
	url := fmt.Sprintf("%s/pixel", baseURL)

	data, _ := json.Marshal(options)
	resp, err := r.makeRequest(ctx, "POST", url, data)
	if err != nil {
		return "", err
	}

	pixel, ok := resp["sensor"].(string)
	if !ok {
		return "", errors.New("failed to parse pixel data")
	}

	return pixel, nil
}

// GenerateSecCptAnswers generates sec-cpt answers.
func (r *RooLink) GenerateSecCptAnswers(ctx context.Context, options *CptChallenge) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/sec-cpt", baseURL)
	data, _ := json.Marshal(options)

	return r.makeRequest(ctx, "POST", url, data)
}

// makeRequest makes an HTTP request and parses the JSON response.
func (r *RooLink) makeRequest(ctx context.Context, method, url string, body []byte, headers ...map[string]string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Default headers
	req.Header.Set("x-api-key", r.APIKey)
	req.Header.Set("Content-Type", "application/json")
	for _, h := range headers {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
