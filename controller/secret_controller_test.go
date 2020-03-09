package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	secretJSON = `{"name": "foo", "namespace": "default", "checksum": "xxxxaaaaqqqq", "data": { "foo":"bar" }}`
)

func TestGetSecret(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/secretreceiver/v1/secret/:namespace/:name")
	c.SetParamNames("namespace", "name")
	c.SetParamValues("default", "foobar")

	// Assertions
	if assert.NoError(t, CheckSecret(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}

func TestPostSlackEvents(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(secretJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/secretreceiver/v1/secret")

	// Assertions
	if assert.NoError(t, CreateSecret(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}

func TestPutSlackEvents(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(secretJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/secretreceiver/v1/secret")

	// Assertions
	if assert.NoError(t, UpdateSecret(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}

func TestDeleteSecret(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/secretreceiver/v1/secret/:namespace/:name")
	c.SetParamNames("namespace", "name")
	c.SetParamValues("default", "foobar")

	// Assertions
	if assert.NoError(t, CheckSecret(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}
