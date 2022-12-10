package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
}

type Config struct {
	ListenAddr string
}

type Server struct {
	Config

	handler  Handler
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]*Peer
	addPeer  chan *Peer
	msgChan  chan *Message
}

func NewServer(cfg Config) *Server {
	return &Server{
		handler: &DefaultHandler{},
		Config:  cfg,
		peers:   make(map[net.Addr]*Peer),
		addPeer: make(chan *Peer),
		msgChan: make(chan *Message),
	}
}

func (s *Server) Start() {
	go s.loop()

	if err := s.listen(); err != nil {
		panic(err)
	}

	fmt.Printf("game server running on port: %s \n", s.ListenAddr)
	s.acceptLoop()
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}

		peer := &Peer{
			conn: conn,
		}

		s.addPeer <- peer

		err = peer.Send([]byte("GGPOKER V0.1-alfa"))
		if err != nil {
			fmt.Println(err)
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		s.msgChan <- &Message{
			From:    conn.RemoteAddr(),
			Payload: bytes.NewReader(buf[:n]),
		}
	}
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}

	s.listener = ln
	return nil
}

func (s *Server) loop() {
	for { //nolint: gosimple
		select {
		case peer := <-s.addPeer:
			s.peers[peer.conn.RemoteAddr()] = peer
			fmt.Printf("new player connected %s\n", peer.conn.RemoteAddr())
		case msg := <-s.msgChan:
			if err := s.handler.HandleMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}
