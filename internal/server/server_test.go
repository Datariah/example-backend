package server_test

import (
	"context"
	"encoding/json"
	"fmt"
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

	expected := map[string]string{
		"message": "Hello!",
	}

	response := map[string]string{}

	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Error while running test: %v", err))
	}

	assert.Equal(t, response["message"], expected["message"])
}
