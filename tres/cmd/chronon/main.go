package main

import (
	"log"
	"net/http"

	"earthcube.org/Project418/garden/tres/internal/mocktime"
	restful "github.com/emicklei/go-restful"
	"github.com/gorilla/mux"
)

// MyServer is the Gorilla mux router struct
type MyServer struct {
	r *mux.Router
}

func main() {
	// Web content
	staticRouter := mux.NewRouter()
	staticRouter.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web"))))
	http.Handle("/", &MyServer{staticRouter})

	// Services
	wsContainer := restful.NewContainer()

	// CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// Add the services
	wsContainer.Add(mocktime.New()) // text search services

	// Running Second Server
	go func() {
		log.Printf("RESTful services on localhost:6789")
		server := &http.Server{Addr: ":6789", Handler: wsContainer}
		server.ListenAndServe()
	}()

	log.Printf("Web server on http://127.0.0.1:7777/")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Let the Gorilla work
	s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
