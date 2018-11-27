package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	port string
	DB   *xorm.Engine
}

// NewServer Constractor
func NewServer() Server {
	return Server{}
}

// Init all vars
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing server ...")
	s.port = ":" + port
	s.DB = db
}

// HomeHandler Handler for Home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on port", s.port)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	http.ListenAndServe(s.port, router)
}
