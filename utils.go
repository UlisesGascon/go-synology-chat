package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var implementedMethods = []string{"chatbot", "user_list", "channel_list"}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func GetUrlByMethod(baseUrl, token string) func(string) (string, error) {
	return func(method string) (string, error) {
		if baseUrl == "" || token == "" {
			return "", errors.New("The baseUrl or the token are missing. Check out the documentation!")
		}

		if !contains(implementedMethods, method) {
			return "", errors.New("Invalid method!")
		}

		return fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Chat.External&method=%s&version=2&token=%%22%s%%22", baseUrl, method, token), nil
	}
}

func MakeGetRequest(url string, ignoreSSLErrors bool) func() (map[string]interface{}, error) {
	return func() (map[string]interface{}, error) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: ignoreSSLErrors},
		}
		client := &http.Client{Transport: tr}

		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

func MakePostRequest(urlStr string, payload string, ignoreSSLErrors bool) (map[string]interface{}, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: ignoreSSLErrors},
	}
	client := &http.Client{Transport: tr}

	formData := url.Values{}
	formData.Set("payload", payload)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GeneratePayloadContent(userIDs []int, text string, file_url ...string) (string, error) {
	payload := map[string]interface{}{
		"user_ids": userIDs,
		"text":     text,
	}
	if len(file_url) > 0 && file_url[0] != "" {
		payload["file_url"] = file_url[0]
	}
	payloadContent, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(payloadContent), nil
}
