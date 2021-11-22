package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var diagramFilePath = "./data/diagram.xml"

func (srv *Server) saveDiagramHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		err := os.Truncate(diagramFilePath, 0)

		if err == nil {
			// TODO: Not necessary with `FormValue`?
			err := r.ParseForm()

			if err == nil {
				// Extract the 'diagram' value from the POST body.
				data := r.FormValue("diagram")

				fmt.Printf("filename: %s\n", diagramFilePath)
				fmt.Printf("data: %s\n", data)

				// Save body to file.
				err := os.WriteFile(diagramFilePath, []byte(fmt.Sprintln(data)), 0644)

				if err == nil {
					fmt.Fprintln(w, `{"status":"success","message":"Success."}`)
				} else {
					fmt.Printf("error: %v\n", err)
					fmt.Fprintf(w, `{"status":"error","message":"Failed writing file: %v"}\n`, err)
				}
			} else {
				fmt.Fprintln(w, `{"status":"error","message":"Failed parsing request body."}`)
			}
		} else {
			fmt.Println("error, failed truncating file")
			fmt.Fprintln(w, `{"status":"error","message":"Failed truncating file."}`)
		}
	} else {
		fmt.Println("error, not post")
		fmt.Fprintln(w, `{"status":"error","message":"This endpoint only accepts POST requests."}`)
	}

}

func (srv *Server) diagramXmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-drawio")

	_, err := os.Stat(diagramFilePath)

	if err == nil {
		// File exists, read it.
		data, err := ioutil.ReadFile(diagramFilePath)

		if err == nil {
			// File read, write it to the client.
			w.Write(data)
		}
	}
}

func (srv *Server) apiHandler(w http.ResponseWriter, r *http.Request) {
	/*
		hosts, err := srv.zbx.GetHosts()

		if err != nil {
			fmt.Fprintf(w, "Error: %v\n", err)
		}

		test := hosts[0]
	*/

	fmt.Fprintf(w, "%s\n", "hi")
}
