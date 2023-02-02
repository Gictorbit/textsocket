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
	defer cli.Stop()
	for {
		msgType, err := cli.ReadOption()
		if err != nil {
			log.Println("read option error")
		}
		switch msgType {
		case api.MessageTypeReverseString:
			if e := cli.SendTextReceiveResult(api.MessageTypeReverseString); e != nil {
				log.Println("reverse string error:", e)
			}
		case api.MessageTypeLowerCaseString:
			if e := cli.SendTextReceiveResult(api.MessageTypeLowerCaseString); e != nil {
				log.Println("lower case string error:", e)
			}
		case api.MessageTypeUpperCaseString:
			if e := cli.SendTextReceiveResult(api.MessageTypeUpperCaseString); e != nil {
				log.Println("upper case string error:", e)
			}
		case api.MessageTypeCountString:
			if e := cli.CountString(api.MessageTypeCountString); e != nil {
				log.Println("count string error:", e)
			}
		}
	}
}
