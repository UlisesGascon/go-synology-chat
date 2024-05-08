package main

import (
    "testing"
	"fmt"
)

func TestContains(t *testing.T) {
    tests := []struct {
        slice []string
        item  string
        want  bool
    }{
        {implementedMethods, "chatbot", true},
        {implementedMethods, "user_list", true},
        {implementedMethods, "channel_list", true},
        {implementedMethods, "non_existent", false},
    }

    for _, tt := range tests {
        if got := contains(tt.slice, tt.item); got != tt.want {
            t.Errorf("contains(%v, %q) = %v, want %v", tt.slice, tt.item, got, tt.want)
        }
    }
}

func TestGetUrlByMethod(t *testing.T) {
    baseUrl := "http://example.com"
    token := "testtoken"

    tests := []struct {
        method string
        want   string
        err    bool
    }{
        {"chatbot", fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Chat.External&method=chatbot&version=2&token=%%22%s%%22", baseUrl, token), false},
        {"user_list", fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Chat.External&method=user_list&version=2&token=%%22%s%%22", baseUrl, token), false},
        {"channel_list", fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Chat.External&method=channel_list&version=2&token=%%22%s%%22", baseUrl, token), false},
        {"non_existent", "", true},
    }

    getUrl := GetUrlByMethod(baseUrl, token)

    for _, tt := range tests {
        got, err := getUrl(tt.method)
        if (err != nil) != tt.err {
            t.Errorf("GetUrlByMethod() error = %v, wantErr %v", err, tt.err)
            return
        }
        if got != tt.want {
            t.Errorf("GetUrlByMethod() = %v, want %v", got, tt.want)
        }
    }
}