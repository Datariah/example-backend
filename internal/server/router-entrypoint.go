package server

import (
	"github.com/Datariah/example-backend/internal/router"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Serve is responsible for defining the routing and starting the negroni server.
func Serve() {
	muxRouter := mux.NewRouter()
	apiV1 := muxRouter.PathPrefix("/api/v1").Subrouter()

	apiV1.HandleFunc("/hello", router.DispatchHello).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(muxRouter)
	n.Run(":80")
}
