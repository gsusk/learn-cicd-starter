package auth_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr error
	}{
		"empty header":     {input: http.Header{}, want: "", wantErr: auth.ErrNoAuthHeaderIncluded},
		"malformed header": {input: http.Header{"Authorization": {"AipKey 12123"}}, want: "", wantErr: errors.New("malformed authorization header")},
		"simple":           {input: http.Header{"Authorization": {"ApiKey D23Si23st4"}}, want: "D23Si23st4", wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := auth.GetAPIKey(tc.input)

			if got == tc.want {
				t.Fatalf("want: %#v, got: %#v", tc.want, got)
			}

			if (err != nil || tc.wantErr != nil) && err.Error() != tc.wantErr.Error() {
				t.Fatalf("want: %#v, got: %#v", tc.wantErr, err)
			}
		})
	}
}
