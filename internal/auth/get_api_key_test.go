package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}

	headers.Set("Authorization", "ApiKey my-secret-api-key-12345")

	gotKey, err := GetAPIKey(headers)
	wantKey := "my-secret-api-key-12345"

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if gotKey != wantKey {
		t.Errorf("got %q, want %q", gotKey, wantKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
    // Create empty headers (do NOT set the Authorization header)
    headers := http.Header{}

    // Call the function
    _, err := GetAPIKey(headers)

    // We EXPECT an error here!
    // Specifically, we expect ErrNoAuthHeaderIncluded
    if err == nil {
        t.Fatalf("expected an error, got nil")
    }

    // You can also verify it's the CORRECT error
    if err != ErrNoAuthHeaderIncluded {
        t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
    }
}