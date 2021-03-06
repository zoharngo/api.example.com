package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"

	"api.example.com/src/router"
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

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on port", s.port)
	r := router.NewRouter()
	r.Init()
	http.ListenAndServe(s.port, r.Router)
}
