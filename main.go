package main

import (
	"github.com/ovguschin90/ggpoker/p2p"
)

func main() {
	cfg := p2p.Config{
		ListenAddr: ":3000",
	}
	srv := p2p.NewServer(cfg)
	srv.Start()
}
