package server

import (
	"fmt"
	"net/http"
)

func (srv *Server) frontendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "frontendHandler")
}

func (srv *Server) apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "apiHandler")
}
