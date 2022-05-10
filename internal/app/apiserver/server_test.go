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

func TestServer_HandleDeleteGroup(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name            string
		payload         interface{}
		expectedCode    int
		createGroup     bool
		createGroupData interface{}
	}{
		{
			name: "valid group create",
			payload: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "3secret",
			},
			expectedCode: http.StatusOK,
			createGroup:  true,
			createGroupData: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "3secret",
			},
		},
		{
			name:         "invalid data",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
			createGroup:  false,
		},
		{
			name: "invalid group owner",
			payload: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "3secret",
			},
			expectedCode: http.StatusBadRequest,
			createGroup:  true,
			createGroupData: map[string]interface{}{
				"group_owner": 3,
				"group_name":  "3secret",
			},
		},
		{
			name: "invalid group name ",
			payload: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "32secret",
			},
			expectedCode: http.StatusBadRequest,
			createGroup:  true,
			createGroupData: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "32secret_unknown group",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.createGroup {
				b := &bytes.Buffer{}
				json.NewEncoder(b).Encode(tc.createGroupData)
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodPost, "/create_group", b)
				s.ServerHTTP(rec, req)
			}
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/delete_group", b)
			s.ServerHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleCreateGroup(t *testing.T) {
	s := newServer(teststore.New(), nil)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid group create",
			payload: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "3secret",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid create group data",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid create group group name include",
			payload: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "3secret",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "valid create group",
			payload: map[string]interface{}{
				"group_owner": 1,
				"group_name":  "32secret",
			},
			expectedCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/create_group", b)
			s.ServerHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

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
