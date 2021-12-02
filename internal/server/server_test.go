package server_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Datariah/example-backend/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestDispatchHello(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/", nil)

	router.DispatchHello(rec, req.WithContext(context.TODO()))

	expected, _ := json.Marshal(&map[string]string{"message": "Hello!"})
	assert.Equal(t, rec.Body.String(), string(expected))
}
