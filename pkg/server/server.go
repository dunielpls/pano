package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dunielpls/pano/pkg/zabbix"
	"github.com/spf13/viper"
)

func New() Server {
	return Server{
		// Return default logger. TOOD: Use klog.
		log: log.Default(),
		zbx: zabbix.New(),
		mux: http.NewServeMux(),
	}
}

func (srv *Server) Start() error {
	srv.initInterrupt()
	srv.initRoutes()
	srv.initZabbix()

	err := http.ListenAndServe(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("server.bind"),
			viper.GetInt("server.port"),
		),
		srv.mux,
	)

	if err != nil {
		fmt.Printf("HTTP server error: %v\n", err)
	}

	return nil
}

func (srv *Server) Block() {
	time.Sleep(1 * time.Second)
}

func (srv *Server) Stop() error {
	return nil
}

func (srv *Server) initRoutes() {
	// Static.
	viewRoute := fmt.Sprintf(
		"%s/",
		strings.TrimRight(
			viper.GetString("server.routes.view"),
			"/",
		),
	)
	editRoute := fmt.Sprintf(
		"%s/",
		strings.TrimRight(
			viper.GetString("server.routes.edit"),
			"/",
		),
	)

	// Serve the same fromtend for both view and edit.
	srv.mux.Handle(viewRoute, http.FileServer(http.Dir("./static")))
	srv.mux.Handle(editRoute, http.FileServer(http.Dir("./static")))

	// Local function to prefix paths with the API prefix.
	// Example: "/hosts" -> "/api/v1/hosts"
	prefixAPI := func(s string) string {
		return fmt.Sprintf(
			"%s%s",
			strings.TrimRight(
				viper.GetString("server.routes.api_prefix"),
				"/",
			),
			s,
		)
	}

	// API.
	srv.mux.HandleFunc(prefixAPI("/hosts/list"), srv.apiHandler)
	srv.mux.HandleFunc(prefixAPI("/hosts/interfaces"), srv.apiHandler)
}

func (srv *Server) initZabbix() {
}

func (srv *Server) BackendIsUp() bool {
	return true
}

func (srv *Server) FrontendIsUp() bool {
	return true
}

func (srv *Server) ZabbixIsUp() bool {
	return true
}
