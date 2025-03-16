//nolint:unused
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
	"github.com/gogf/gf/v2/encoding/gjson"
)

var (
	baseURL = "http://localhost:8080"
)

type apiFeature struct {
	httpClient *http.Client
	req        *http.Request
	resp       *http.Response
	logger     *slog.Logger
	token      string
}

func (a *apiFeature) reset(*godog.Scenario) error {
	a.req = nil
	a.resp = nil
	a.logger = slog.Default()
	return nil
}

func (a *apiFeature) iAuthenticateWithUsernameAndPassword(username, password string) (err error) {
	content := strings.NewReader(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))

	if err = a.httpCall(http.MethodPost, "/v1/auth", content, nil); err != nil {
		return err
	}

	defer a.resp.Body.Close()

	bodyBytes, err := io.ReadAll(a.resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read authentication response body: %v", err)
	}

	var j *gjson.Json
	if j, err = gjson.DecodeToJson(bodyBytes); err != nil {
		return fmt.Errorf("unable to unmarshal authentication response: %v", err)
	}

	a.logger.Info("j:", slog.Any("j", j))

	a.token = j.Get("data.access_token").String()

	return nil
}

func (a *apiFeature) iSendRequestTo(method, endpoint string) error {
	return a.httpCall(method, endpoint, nil, nil)
}

func (a *apiFeature) iSendRequestToWithPayload(method, endpoint string, body *godog.DocString) error {
	reader := strings.NewReader(body.Content)
	return a.httpCall(method, endpoint, reader, nil)
}

func (a *apiFeature) httpCall(method, endpoint string, content io.Reader, writer *multipart.Writer) error {
	url := baseURL + endpoint

	req, err := http.NewRequest(method, url, content)
	if err != nil {
		return fmt.Errorf("unable to prepare http request: %w", err)
	}

	if a.token != "" {
		req.Header.Add("Authorization", "Bearer "+a.token)
	}

	if writer != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("unable to query http request: %w", err)
	}

	a.req = req
	a.resp = resp

	return nil
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if a.resp == nil {
		return fmt.Errorf("http response is nil")
	}

	if code != a.resp.StatusCode {
		if a.resp.StatusCode >= 400 {
			bodyBytes, err := io.ReadAll(a.resp.Body)
			if err != nil {
				return fmt.Errorf("unable to read request body: %v", err)
			}

			return fmt.Errorf("expected response code to be: %d, but actual is: %d, response message: %s", code, a.resp.StatusCode, string(bodyBytes))
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.StatusCode)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON3(body *godog.DocString) (err error) {
	if a.resp == nil {
		return fmt.Errorf("http response is nil")
	}

	var expected, actual map[string]any

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	bodyBytes, err := io.ReadAll(a.resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}

	var j *gjson.Json
	if j, err = gjson.DecodeToJson(bodyBytes); err != nil {
		return fmt.Errorf("unable to unmarshal authentication response: %v", err)
	}

	actual = j.Get("data").Map()

	sortArray(expected)
	sortArray(actual)

	jsonExpected, _ := json.MarshalIndent(expected, "", " ")
	jsonActual, _ := json.MarshalIndent(actual, "", " ")

	// the matching may be adapted per different requirementa.
	if !bytes.Equal(jsonExpected, jsonActual) {
		return fmt.Errorf("expected JSON does not match.\n-> expected:\n%v\n-> actual:\n%v", string(jsonExpected), string(jsonActual))
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON2(body *godog.DocString) (err error) {
	if a.resp == nil {
		return fmt.Errorf("http response is nil")
	}

	var expected, actual map[string]any

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	bodyBytes, err := io.ReadAll(a.resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}

	var j *gjson.Json
	if j, err = gjson.DecodeToJson(bodyBytes); err != nil {
		return fmt.Errorf("unable to unmarshal authentication response: %v", err)
	}

	actual = j.Get("data").Map()

	sortArray(expected)
	sortArray(actual)

	jsonExpected, _ := json.MarshalIndent(expected, "", " ")
	jsonActual, _ := json.MarshalIndent(actual, "", " ")

	// the matching may be adapted per different requirementa.
	if !bytes.Equal(jsonExpected, jsonActual) {
		return fmt.Errorf("expected JSON does not match.\n-> expected:\n%v\n-> actual:\n%v", string(jsonExpected), string(jsonActual))
	}

	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	if a.resp == nil {
		return fmt.Errorf("http response is nil")
	}

	var expected, actual map[string]any

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	bodyBytes, err := io.ReadAll(a.resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}

	var j *gjson.Json
	if j, err = gjson.DecodeToJson(bodyBytes); err != nil {
		return fmt.Errorf("unable to unmarshal authentication response: %v", err)
	}

	actual = j.Get("data").Map()

	sortArray(expected)
	sortArray(actual)

	jsonExpected, _ := json.MarshalIndent(expected, "", " ")
	jsonActual, _ := json.MarshalIndent(actual, "", " ")

	// the matching may be adapted per different requirementa.
	if !bytes.Equal(jsonExpected, jsonActual) {
		return fmt.Errorf("expected JSON does not match.\n-> expected:\n%v\n-> actual:\n%v", string(jsonExpected), string(jsonActual))
	}
	return nil
}

func sortArray(data map[string]interface{}) {
	for _, v := range data {
		switch vv := v.(type) {
		case []interface{}:
			sortArray(vv[0].(map[string]interface{}))
		case map[string]interface{}:
			sortArray(vv)
		}
	}
}
