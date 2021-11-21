package server

import (
	"fmt"
	"os/exec"
	"strings"

	color "github.com/fatih/color"
)

func (srv *Server) CLIVersion() string {
	// Logic for `pano version`.
	commit := "unknown"

	stdout, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()

	if err == nil {
		commit = strings.TrimSpace(string(stdout))
	}

	return fmt.Sprintf("Pano version 0.1 (commit %s)\n", commit)
}

func (srv *Server) CLIStatus() string {
	// Logic for `pano status`.
	backend_status := color.New(color.FgRed).Sprint("down")
	if srv.BackendIsUp() {
		backend_status = color.New(color.FgGreen).Sprint("up")
	}

	frontend_status := color.New(color.FgRed).Sprint("down")
	if srv.FrontendIsUp() {
		frontend_status = color.New(color.FgGreen).Sprint("up")
	}

	zabbix_status := color.New(color.FgRed).Sprint("down")
	if srv.ZabbixIsUp() {
		zabbix_status = color.New(color.FgGreen).Sprint("up")
	}

	out := []string{
		fmt.Sprintf("Backend:  %s\n", backend_status),
		fmt.Sprintf("Frontend: %s\n", frontend_status),
		fmt.Sprintf("Zabbix:   %s\n", zabbix_status),
	}

	return strings.Join(out, "")
}

func (srv *Server) CLIStart() error {
	// Logic for `pano start`.
	return srv.Start()
}

func (srv *Server) CLIStop() error {
	// Logic for `pano stop`.
	return srv.Stop()
}
