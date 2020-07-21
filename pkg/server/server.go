package server

import (
	"net"
	"time"
)

//Server is the core CacheDb struct
//Server has the tcp listener, channels to handle Signals and map of client connections
type Server struct {
	listener    net.Listener
	exit        chan struct{}
	quit        chan struct{}
	connections map[int]net.Conn
	timeout     time.Duration
	db          CacheDB
}

//NewServer creates a new DB server and returns that instance to main func
func NewServer() *Server {
	return nil
}

//Stop stops the CacheDB server
func (s *Server) Stop() {

}
