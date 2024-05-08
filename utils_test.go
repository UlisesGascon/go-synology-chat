package synologychat

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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

func TestMakeGetRequest(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if req.URL.String() != "/test" {
			t.Errorf("want: '/test', got: '%s'", req.URL.String())
		}
		// Send response to be tested
		rw.Write([]byte(`{"hello": "world"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use server.URL with MakeGetRequest to test
	result, err := MakeGetRequest(server.URL+"/test", false)()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check the result
	if result["hello"] != "world" {
		t.Errorf("want: 'world', got: '%s'", result["hello"])
	}
}

func TestMakeGetRequest_SSL(t *testing.T) {
	// Start a local HTTPS server
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`{"hello": "world"}`))
	}))
	// Disable logging
	server.Config.ErrorLog = log.New(io.Discard, "", 0)
	// Close the server when test finishes
	defer server.Close()

	// Use server.URL with MakeGetRequest to test
	_, err := MakeGetRequest(server.URL+"/test", false)()
	if err == nil {
		t.Errorf("expected SSL error, got nil")
	}

	// Now try with ignoring SSL errors
	_, err = MakeGetRequest(server.URL+"/test", true)()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGeneratePayloadContent(t *testing.T) {
	userIDs := []int{1, 2, 3}
	text := "Hello, World!"
	file_url := "http://example.com/file"

	// Test without file_url
	result, err := GeneratePayloadContent(userIDs, text)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"text":"Hello, World!","user_ids":[1,2,3]}`
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test with file_url
	result, err = GeneratePayloadContent(userIDs, text, file_url)
	if err != nil {
		t.Fatal(err)
	}
	expected = `{"file_url":"http://example.com/file","text":"Hello, World!","user_ids":[1,2,3]}`
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMakePostRequest(t *testing.T) {
	// Create a test server that responds with a predefined JSON string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "success", "message": "Hello, World!"}`))
	}))
	defer server.Close()

	// Call the function with the test server's URL
	result, err := MakePostRequest(server.URL, "payload", false)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the result
	expected := map[string]interface{}{
		"status":  "success",
		"message": "Hello, World!",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}
