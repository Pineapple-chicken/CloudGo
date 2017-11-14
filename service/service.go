package service

import (
	"net/http"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initRouter(r *mux.Router, formatter *render.Render) {
	r.HandleFunc("/hello/{name}", helloHandler(formatter)).Methods("GET")
}

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options {
	    IndentJSON: true,
	})
	// New mux and server
	r := mux.NewRouter()
	server := negroni.Classic()
	// Add router
	initRouter(r, formatter)
	server.UseHandler(r)
	return server
}

func helloHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		formatter.JSON(w, http.StatusOK, map[string]string{"name": name})
	}
}
