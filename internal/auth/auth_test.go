package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "No Authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "Malformed Authorization header (missing ApiKey)",
			headers:     http.Header{"Authorization": []string{"Bearer token"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name:        "Malformed Authorization header (incorrect scheme)",
			headers:     http.Header{"Authorization": []string{"ApiKey123 token"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name:        "Valid Authorization header",
			headers:     http.Header{"Authorization": []string{"ApiKey 1234567890"}},
			expectedKey: "1234567890",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			// Check if the error matches the expected error
			if !errors.Is(err, tt.expectedErr) && (err == nil || tt.expectedErr == nil || err.Error() != tt.expectedErr.Error()) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}

			// Check if the returned key matches the expected key
			if key != tt.expectedKey {
				t.Errorf("expected API key %s, got %s", tt.expectedKey, key)
			}
		})
	}

}
