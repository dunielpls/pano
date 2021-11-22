package server

import (
	"log"
	"net/http"

	"github.com/dunielpls/pano/pkg/zabbix"
)

type Server struct {
	log *log.Logger
	zbx *zabbix.Zabbix
	mux *http.ServeMux
}

type Host struct {
	Id          int
	Name        string
	DisplayName string
	Groups      []int
}
