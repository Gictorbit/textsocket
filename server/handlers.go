package server

import (
	"encoding/json"
	pc "github.com/Gictorbit/textsocket/api"
	"net"
	"strings"
)

// ReverseString reverse a string
func (s *Server) ReverseString(req *pc.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	resultStr := []byte(input)
	for i, j := 0, len(resultStr)-1; i < j; i, j = i+1, j-1 {
		resultStr[i], resultStr[j] = resultStr[j], resultStr[i]
	}
	response := &pc.PacketBody{
		MessageType: req.MessageType,
		Data:        resultStr,
	}
	return pc.SendPacket(conn, response)
}

// UpperCaseString makes a string upper case
func (s *Server) UpperCaseString(req *pc.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	response := &pc.PacketBody{
		MessageType: req.MessageType,
		Data:        []byte(strings.ToUpper(input)),
	}
	return pc.SendPacket(conn, response)
}

// LowerCaseString makes a string lower case
func (s *Server) LowerCaseString(req *pc.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	response := &pc.PacketBody{
		MessageType: req.MessageType,
		Data:        []byte(strings.ToLower(input)),
	}
	return pc.SendPacket(conn, response)
}

// CountString makes a string lower case
func (s *Server) CountString(req *pc.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	words := strings.Split(input, " ")
	letters := 0
	for _, w := range words {
		letters += len(w)
	}
	result := pc.CountString{
		Words:   len(words),
		Letters: letters,
	}
	resByte, err := json.Marshal(result)
	if err != nil {
		return err
	}
	response := &pc.PacketBody{
		MessageType: req.MessageType,
		Data:        resByte,
	}
	return pc.SendPacket(conn, response)
}
