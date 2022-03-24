package gorillamux

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGorillaMuxHandler(t *testing.T) {
	// Testcase Setup
	expectedNamespace := `this/is/namespace`
	expectedType := `thisistype`
	// Setup serverside
	handler := NewHandler()
	// Make the client-side request
	request, _ := http.NewRequest("GET", fmt.Sprintf("https://example.com/api/%s/%s/versions", expectedNamespace, expectedType), nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	// Parse the response, as we would client-side
	responseBytes, err := ioutil.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)
	lines := strings.Split(string(responseBytes), "\n")
	// Testcase Validation
	assert.Contains(t, lines, expectedNamespace)
	assert.Contains(t, lines, expectedType)
}
