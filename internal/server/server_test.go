package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Datariah/example-backend/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	router.DispatchHello(rec, req)
	expected, _ := json.Marshal(&map[string]string{"message": "Hello!"})
	assert.Equal(t, rec.Body.String(), string(expected))
}
