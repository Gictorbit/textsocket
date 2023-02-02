package server

import (
	pc "github.com/Gictorbit/textsocket/api"
	"log"
	"net"
	"sync"
)

type Empty struct{}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitChan   chan Empty
	wg         sync.WaitGroup
	log        *log.Logger
}

type ServerInterface interface {
	Start()
	Stop()
	ReverseString(req *pc.PacketBody, conn net.Conn) error
	UpperCaseString(req *pc.PacketBody, conn net.Conn) error
	LowerCaseString(req *pc.PacketBody, conn net.Conn) error
	CountString(req *pc.PacketBody, conn net.Conn) error
}

var (
	_ ServerInterface = &Server{}
)

func NewServer(listenAddr string) ServerInterface {
	return &Server{
		listenAddr: listenAddr,
		quitChan:   make(chan Empty),
		wg:         sync.WaitGroup{},
		log:        log.Default(),
	}
}

func (s *Server) Start() {
	ln, err := net.Listen(pc.SocketType, s.listenAddr)
	if err != nil {
		s.log.Println("failed to listen: ", err.Error())
		return
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptConnections()
	s.log.Println("server started", "ListenAddress: "+s.listenAddr)
	<-s.quitChan
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			s.log.Println("accept connection error", err.Error())
			continue
		}
		s.log.Println("new Connection to the server", "Address: "+conn.RemoteAddr().String())
		s.wg.Add(1)
		go s.HandleConnection(conn)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	defer s.wg.Done()
	packet, err := pc.ReadPacket(conn)
	if err != nil {
		s.log.Println("error read packet", err)
	}
	switch packet.MessageType {
	case pc.MessageTypeReverseString:
		if e := s.ReverseString(packet, conn); e != nil {
			s.log.Println("reverse string failed", "Address: "+conn.RemoteAddr().String())
		}
	case pc.MessageTypeCountString:
		if e := s.CountString(packet, conn); e != nil {
			s.log.Println("count string failed", "Address: "+conn.RemoteAddr().String())
		}
	case pc.MessageTypeUpperCaseString:
		if e := s.UpperCaseString(packet, conn); e != nil {
			s.log.Println("upper case string failed", "Address: "+conn.RemoteAddr().String())
		}
	case pc.MessageTypeLowerCaseString:
		if e := s.LowerCaseString(packet, conn); e != nil {
			s.log.Println("lower case string failed", "Address: "+conn.RemoteAddr().String())
		}
	}

}

func (s *Server) Stop() {
	s.wg.Wait()
	s.quitChan <- Empty{}
	s.log.Println("stop server")
}
