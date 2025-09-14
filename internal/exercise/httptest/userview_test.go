package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	users := map[string]User{
		"u1": {
			ID:        "u1",
			FirstName: "Misha",
			LastName:  "Popov",
		},
		"u2": {
			ID:        "u2",
			FirstName: "Sasha",
			LastName:  "Popov",
		},
	}

	type want struct {
		code     int
		response string
	}
	tests := []struct {
		name    string
		user_id string
		want    want
	}{
		{"u1", "u1", want{200, `{"ID":"u1","FirstName":"Misha","LastName":"Popov"}`}},
		{"u2", "u2", want{200, `{"ID":"u2","FirstName":"Sasha","LastName":"Popov"}`}},
		{"u_empty", "", want{400, ""}},
		{"u_not_found", "u3", want{404, ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/userview?user_id="+tt.user_id, nil)
			w := httptest.NewRecorder()
			handler := UserViewHandler(users)
			handler(w, request)
			res := w.Result()
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, tt.want.code, res.StatusCode)
			if res.StatusCode == 200 {
				assert.JSONEq(t, tt.want.response, string(resBody))
			}
		})
	}
}
