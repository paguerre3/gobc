package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Create a test Echo context
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Call the Ping function
	err = Ping(c)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the response
	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "pong", response["message"])
}
