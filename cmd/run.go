package main

import (
	"fmt"
	pc "github.com/Gictorbit/textsocket/api"
	"github.com/Gictorbit/textsocket/server"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "server":
		RunServer()
	}
}

func RunServer() {
	listenAddr := net.JoinHostPort(pc.HostAddress, fmt.Sprintf("%d", pc.PortNumber))
	srv := server.NewServer(listenAddr)
	go srv.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	srv.Stop()
}
