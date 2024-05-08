package main

import (
	"errors"
	"fmt"
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
