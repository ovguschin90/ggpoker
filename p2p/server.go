package p2p

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddr  string
	Version     string
	GameVariant GameVariant
}

type Server struct {
	Config

	transport *TCPTransport
	peers     map[net.Addr]*Peer
	addPeer   chan *Peer
	delPeer   chan *Peer
	msgChan   chan *Message
}

func NewServer(cfg Config) *Server {
	s := &Server{
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

	logrus.WithFields(logrus.Fields{
		"port": s.ListenAddr,
		"type": s.GameVariant,
	}).Info("started new game server")

	s.transport.ListenAndAccept()
}

func (s *Server) SendHandshake(p *Peer) error {
	msg, err := Encode(s, p)
	if err != nil {
		return err
	}

	return p.Send(msg)
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
			go s.SendHandshake(peer)

			if err := s.handshake(peer); err != nil {
				logrus.Error("handshake failed: ", err)
				continue
			}

			//TODO: check max players and other game state logic
			go peer.ReadLoop(s.msgChan)

			logrus.WithFields(logrus.Fields{
				"addr": peer.conn.RemoteAddr(),
			}).Info("new player connected")

			s.peers[peer.conn.RemoteAddr()] = peer
		case msg := <-s.msgChan:
			if err := s.handleMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}

func (s *Server) handshake(p *Peer) error {
	hs, err := Decode(p)

	if err != nil {
		return err
	}

	fmt.Printf("hs =>  %+v\n", hs)
	logrus.WithFields(logrus.Fields{
		"peer":    p.conn.RemoteAddr(),
		"version": hs.Version,
		"variant": hs.GameVariant,
	}).Info("recieved handshake")

	return nil
}

func (s *Server) handleMessage(msg *Message) error {
	fmt.Printf("%+v\n", msg)
	return nil
}
