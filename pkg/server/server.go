package server

import (
	"fmt"
	"log"
	"net/http"
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
	// Frontend.
	srv.mux.Handle("/", http.FileServer(http.Dir("./static")))

	// API.
	srv.mux.HandleFunc("/api/v1/diagram.xml", srv.diagramXmlHandler)
	srv.mux.HandleFunc("/api/v1/save_diagram", srv.saveDiagramHandler)
	srv.mux.HandleFunc("/api/v1/hosts/list", srv.apiHandler)
	srv.mux.HandleFunc("/api/v1/hosts/interfaces", srv.apiHandler)
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
