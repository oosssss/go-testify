package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedStatus := http.StatusOK
	codeResult := responseRecorder.Code
	assert.Equal(t, expectedStatus, codeResult)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=dombay", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedStatus := http.StatusBadRequest
	codeResult := responseRecorder.Code
	assert.Equal(t, expectedStatus, codeResult)

	expectedBody := `wrong city value`
	body := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedStatus := http.StatusOK
	codeResult := responseRecorder.Code
	require.Equal(t, expectedStatus, codeResult)

	body := responseRecorder.Body.String()
	cityList := strings.Split(body, ",")

	require.Len(t, cityList, totalCount)
}
