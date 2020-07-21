package server

import (
	"log"
	"net"
	"time"

	"github.com/kanapuli/CacheDB/pkg/db"
)

const (
	port        = "9097"
	connTimeout = 10
)

//Server is the core CacheDb struct
//Server has the tcp listener, channels to handle Signals and map of client connections
type Server struct {
	listener    net.Listener
	exit        chan struct{}
	quit        chan struct{}
	cacheDB     db.CacheDB
	connections map[int]net.Conn
	timeout     time.Duration
}

//NewServer creates a new DB server and returns that instance to main func
func NewServer() *Server {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to create listener on port %v: %v", port, err)
	}
	srv := &Server{
		listener:    l,
		exit:        make(chan struct{}),
		quit:        make(chan struct{}),
		cacheDB:     db.NewCacheDB(),
		connections: make(map[int]net.Conn),
		timeout:     connTimeout * time.Second,
	}
	return srv
}

//Stop stops the CacheDB server
func (s *Server) Stop() {

}
