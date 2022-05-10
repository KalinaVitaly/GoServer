package apiserver

import (
	"Diplom/internal/app/store/teststore"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleDeleteUser(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
		createUser   bool
		userData     interface{}
	}{
		{
			name: "valid delete user",
			payload: map[string]interface{}{
				"email":    "3user@example.org",
				"password": "3secret",
			},
			expectedCode: http.StatusOK,
			createUser:   true,
			userData: map[string]interface{}{
				"email":    "3user@example.org",
				"password": "3secret",
			},
		},
		{
			name:         "invalid delete user (unknown user)",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
			createUser:   false,
		},
		{
			name: "invalid delete user invalid params email",
			payload: map[string]interface{}{
				"email":    "2user@example.org",
				"password": "2secret",
			},
			expectedCode: http.StatusBadRequest,
			createUser:   true,
			userData: map[string]interface{}{
				"email":    "2user@examp1le.org",
				"password": "2secret",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.createUser {
				b := &bytes.Buffer{}
				json.NewEncoder(b).Encode(tc.userData)
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodPost, "/create_users", b)
				s.ServerHTTP(rec, req)
			}
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/delete_user", b)
			s.ServerHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"email":    "user@example.org",
				"password": "secret",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid params",
			payload: map[string]interface{}{
				"email":    "invalid",
				"password": "short",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/create_users", b)
			s.ServerHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
