package main

import (
	"github.com/gorilla/mux"
	"gofacilito/conectarGeneral"
	"gofacilito/login"
	"database/sql"
	"net/http"
	"time"
	"os"
)

var ConexionDB *sql.DB = conectarGeneral.Crearpoll()


func verificarserver(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" || r.Method == "options" {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "authorization,content-type")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.WriteHeader(http.StatusOK)

	} else {
		login.TraerUsuario(w, r, ConexionDB)
	}
}

func main() {
	mex := mux.NewRouter().StrictSlash(true)
	mex.HandleFunc("/hola", verificarserver)
	mex.PathPrefix("/login").Handler(http.StripPrefix("/login", http.FileServer(http.Dir("public"))))
	mex.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("public"))))
	//public := http.FileServer(http.Dir("public"))
	//mex.Handle("/", public)
	port := os.Getenv("PORT")
	if port == "" {
	port = "90"
	}

	myServer := &http.Server{
		Addr:           ":"+port,
		Handler:        mex,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	myServer.ListenAndServe()
}
