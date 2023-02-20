package routers

import (
	"TCP_Packet/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetIPS(router *mux.Router) *mux.Router {
	router.Handle("/update/v1/ips",
		negroni.New(
			negroni.HandlerFunc(controllers.UpdateIPS),
		)).Methods("POST")
	router.Handle("/delete/v1/ips",
		negroni.New(
			negroni.HandlerFunc(controllers.DeleteIPS),
		)).Methods("DELETE")
	router.Handle("/readall/v1/ips",
		negroni.New(
			negroni.HandlerFunc(controllers.GetIPS),
		)).Methods("GET")
	return router

}
