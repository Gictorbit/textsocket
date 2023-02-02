package main

import (
	"fmt"
	"github.com/Gictorbit/textsocket/api"
	"github.com/Gictorbit/textsocket/client"
	"github.com/Gictorbit/textsocket/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "server":
		RunServer()
	case "client":
		RunClient()
	}
}

func RunServer() {
	listenAddr := net.JoinHostPort(api.HostAddress, fmt.Sprintf("%d", api.PortNumber))
	srv := server.NewServer(listenAddr)
	go srv.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	srv.Stop()
}

func RunClient() {
	listenAddr := net.JoinHostPort(api.HostAddress, fmt.Sprintf("%d", api.PortNumber))
	log.Println("server address is ", listenAddr)
	cli := client.NewClient(listenAddr, log.Default())
	if e := cli.Connect(); e != nil {
		log.Fatal(e)
	}

	cli.Stop()
}
