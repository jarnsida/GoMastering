package main

import (
	"fmt"
	"net/http"

	"github.com/jarnsida/GoMastering/HYDRA/HYDRA/hlogger"
	//	"github.com/jarnsida/OzonIMDG_Case/logger"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to Hydra software system")

	logger.Println("Recieved an http Get request on root url")
}
