package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	hdr := http.Header{}
	hdr.Add("Host", "test.com")
	hdr.Add("Authorization", "ApiKey 12345")
	hdr.Add("Content-Type", "application/json")

	tests := []struct {
		name   string
		header http.Header
		key    string
		e      bool
	}{
		{
			name:   "no-header",
			header: nil,
			key:    "",
			e:      true,
		},
		{
			name:   "with-bearer",
			header: hdr,
			key:    "12345",
			e:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)

			if err != nil && !tt.e {
				t.Errorf("unexpected error: %v", err)
			}
			if err == nil && tt.e {
				t.Errorf("did not error when expected")
			}
			if key != tt.key {
				t.Errorf("key: have %s, want %s", key, tt.key)
			}
		})
	}
}
