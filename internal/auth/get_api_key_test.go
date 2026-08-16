package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyNilRequestHeaders(t *testing.T) {
	apiKey, err := GetAPIKey(nil)
	if err == nil {
		t.Errorf("Expected error, got %v", apiKey)
	}
	if apiKey != "" {
		t.Errorf("Expected empty string, got %s", apiKey)
	}
}

func TestGetAPIKey(t *testing.T) {
	// Create sample req to then create headers
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Errorf("Expected no error creating sample request, got %v", err)
	}
	req.Header.Set("Authorization", "ApiKey my_api_key")

	// Test GetAPIKey with headers
	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "my_api_key" {
		t.Errorf("Expected my_api_key, got %s", apiKey)
	}
}

func TestGetAPIKeyBearerToken(t *testing.T) {
	// Create sample req to then create headers
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Errorf("Expected no error creating sample request, got %v", err)
	}
	req.Header.Set("Authorization", "Bearer my_api_key")

	// Test GetAPIKey with headers
	apiKey, err := GetAPIKey(req.Header)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
	}
	if apiKey == "my_api_key" {
		t.Error("Did not expect my_api_key")
	}
}
