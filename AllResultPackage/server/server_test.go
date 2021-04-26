package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mserebryaakov/course-golang/build/server"
)

func TestServer(t *testing.T) {
	testBody := "body"

	request, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(testBody))
	response := httptest.NewRecorder()
	server.Handler(response, request)

	getBody := response.Body.String()
	if getBody != testBody {
		t.Fatal("Incorrect get body")
	}
}
