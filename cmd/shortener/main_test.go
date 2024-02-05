package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_treatURL(t *testing.T) {
	tests := []struct {
		method                    string
		expectedCode              int
		expectedBody              string
		expectedHeaderContentType string
		expectedHeaderLocation    string
	}{
		{method: http.MethodDelete, expectedCode: http.StatusBadRequest, expectedBody: "", expectedHeaderContentType: "", expectedHeaderLocation: ""},
		{method: http.MethodPut, expectedCode: http.StatusBadRequest, expectedBody: "", expectedHeaderContentType: "", expectedHeaderLocation: ""},
		{method: http.MethodPost, expectedCode: http.StatusCreated, expectedBody: "", expectedHeaderContentType: "text/plain", expectedHeaderLocation: ""},
		{method: http.MethodGet, expectedCode: http.StatusTemporaryRedirect, expectedBody: "", expectedHeaderContentType: "", expectedHeaderLocation: ""},
	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()

			treatURL(w, r)

			assert.Equal(t, tt.expectedCode, w.Code, fmt.Sprintf("expected: %d\nactual: %d", tt.expectedCode, w.Code))
			if tt.expectedBody != "" {
				assert.Equal(t, tt.expectedBody, w.Body.String(), fmt.Sprintf("expected: %d\nactual: %d", tt.expectedCode, w.Code))
			}

			if tt.expectedHeaderContentType != "" {
				assert.Equal(
					t,
					tt.expectedHeaderContentType,
					w.Header().Get("Content-Type"),
					fmt.Sprintf(
						"expected: %s\nactual: %s",
						tt.expectedHeaderContentType,
						w.Header().Get("Content-Type"),
					),
				)
			}

			if tt.expectedHeaderLocation != "" {
				assert.Equal(
					t,
					tt.expectedHeaderLocation,
					w.Header().Get("Location"),
					fmt.Sprintf(
						"expected: %s\nactual: %s",
						tt.expectedHeaderLocation,
						w.Header().Get("Location"),
					),
				)
			}
		})
	}
}
