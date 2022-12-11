package p2p

import (
	"fmt"
	"net"
	"sync"

	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddr string
	Verion     string
}

type Server struct {
	Config

	handler   Handler
	transport *TCPTransport
	mu        sync.RWMutex
	peers     map[net.Addr]*Peer
	addPeer   chan *Peer
	delPeer   chan *Peer
	msgChan   chan *Message
}

func NewServer(cfg Config) *Server {
	s := &Server{
		handler: &DefaultHandler{},
		Config:  cfg,
		peers:   make(map[net.Addr]*Peer),
		addPeer: make(chan *Peer),
		delPeer: make(chan *Peer),
		msgChan: make(chan *Message),
	}

	tr := NewTCPtransport(s.ListenAddr)
	s.transport = tr
	tr.AddPeer = s.addPeer
	tr.DelPeer = s.delPeer

	return s
}

func (s *Server) Start() {
	go s.loop()

	fmt.Printf("game server running on port: %s \n", s.ListenAddr)

	s.transport.ListenAndAccept()
}

func (s *Server) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	peer := &Peer{
		conn: conn,
	}

	s.addPeer <- peer

	return peer.Send([]byte("GGPOKER V0.1-alfa"))
}

func (s *Server) loop() {
	for { //nolint: gosimple
		select {
		case peer := <-s.delPeer:
			logrus.WithFields(logrus.Fields{
				"addr": peer.conn.RemoteAddr(),
			}).Info("new player disconnected")
			delete(s.peers, peer.conn.RemoteAddr())
		case peer := <-s.addPeer:
			logrus.WithFields(logrus.Fields{
				"addr": peer.conn.RemoteAddr(),
			}).Info("new player connected")
			s.peers[peer.conn.RemoteAddr()] = peer
		case msg := <-s.msgChan:
			if err := s.handler.HandleMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}
