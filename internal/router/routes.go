package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// DispatchHello triggers a 200 response saying Hello!
func DispatchHello(writer http.ResponseWriter, req *http.Request) {
	renderer := render.New()

	timeBytes, err := json.Marshal(time.Now())
	if err != nil {
		err = renderer.JSON(writer, http.StatusInternalServerError, map[string]string{
			"message":   fmt.Sprintf("Error: %v", err),
			"timestamp": "EMPTY",
		})
		if err != nil {
			log.WithError(err).Error("Error generating timestamp for response")
		}
	}

	err = renderer.JSON(writer, http.StatusOK, map[string]string{
		"message":   "Hello!",
		"timestamp": string(timeBytes),
	})
	// Let's handle a possible error while replying back
	if err != nil {
		log.WithError(err).Error("There was a problem replying back to the client")
	}
}
