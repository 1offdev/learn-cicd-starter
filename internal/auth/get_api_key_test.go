package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "gets API key from authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			want: "abc1234",
		},
		{
			name:    "returns error when authorization header is missing",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name: "returns error when authorization header is malformed",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
