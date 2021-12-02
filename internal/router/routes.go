package router

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// DispatchHello triggers a 200 response saying Hello!
func DispatchHello(w http.ResponseWriter, req *http.Request) {
	r := render.New()

	err := r.JSON(w, http.StatusOK, map[string]string{
		"message": "Hello!",
	})
	// Let's handle a possible error while replying back
	if err != nil {
		log.WithError(err).Error("There was a problem replying back to the client")
	}
}
