package p2p

import (
	"bytes"
	"encoding/gob"
)

type HandshakeMessage struct {
	GameVariant GameVariant
	Version     string
}

func Encode(s *Server, p *Peer) ([]byte, error) {
	hs := &HandshakeMessage{
		GameVariant: s.GameVariant,
		Version:     s.Version,
	}

	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(hs); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(p *Peer) (*HandshakeMessage, error) {
	hs := &HandshakeMessage{}

	if err := gob.NewDecoder(p.conn).Decode(hs); err != nil {
		return hs, err
	}

	return hs, nil
}
