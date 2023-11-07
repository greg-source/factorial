package internal

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	url             = "/calculate"
	reqBodyTemplate = "{\"a\": %d,\"b\": %d\n}"
)

func Test_factorialHandler(t *testing.T) {
	tests := []struct {
		name    string
		reqBody string
		message string
		status  int
	}{
		{
			name:    "Ok",
			reqBody: fmt.Sprintf(reqBodyTemplate, 5, 5),
			message: "{\"a_factorial\":120,\"b_factorial\":120}",
			status:  200,
		},
		{
			name:    "Negative a",
			reqBody: fmt.Sprintf(reqBodyTemplate, -5, 5),
			message: "{\"error\":\"incorrect input\"}",
			status:  400,
		},
		{
			name:    "Negative b",
			reqBody: fmt.Sprintf(reqBodyTemplate, 5, -5),
			message: "{\"error\":\"incorrect input\"}",
			status:  400,
		},
		{
			name:    "a not exist",
			reqBody: fmt.Sprintf("{\"aa\": %d,\"b\": %d\n}", 5, -5),
			message: "{\"error\":\"incorrect input\"}",
			status:  400,
		},
		{
			name:    "a not exist",
			reqBody: fmt.Sprintf("{\"a\": %d,\"bb\": %d\n}", 5, -5),
			message: "{\"error\":\"incorrect input\"}",
			status:  400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := httprouter.New()
			router.Handle("POST", url, factorialHandler)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(tt.reqBody)))
			router.ServeHTTP(w, r)
			assert.Equal(t, tt.message, w.Body.String())
			assert.Equal(t, tt.status, w.Code)
		})
	}
}
