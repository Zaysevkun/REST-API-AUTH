package api

import (
	"bytes"
	"encoding/json"
	"github.com/Zaysevkun/RESTful-API/model"
	"github.com/Zaysevkun/RESTful-API/storage/teststorage"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestServer_HandleUsersCreate table test
func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(teststorage.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name               string
		payload            interface{}
		expectedStatusCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "invalid",
			payload:            "invalid",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedStatusCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	storage := teststorage.New()
	u := model.TestUser(t)
	storage.User().Create(u)
	s := NewServer(storage, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name               string
		payload            interface{}
		expectedStatusCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid",
			payload:            "invalid",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sessions", bytes.NewBuffer(body))
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedStatusCode, rec.Code)
		})
	}
}
