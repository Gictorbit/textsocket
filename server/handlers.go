package server

import (
	"encoding/json"
	"fmt"
	api "github.com/Gictorbit/textsocket/api"
	"net"
	"strings"
)

// ReverseString reverse a string
func (s *Server) ReverseString(req *api.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	resultStr := []byte(input)
	for i, j := 0, len(resultStr)-1; i < j; i, j = i+1, j-1 {
		resultStr[i], resultStr[j] = resultStr[j], resultStr[i]
	}
	response := &api.PacketBody{
		MessageType: req.MessageType,
		Data:        resultStr,
	}
	s.log.Println("string reversed", input, "->", string(resultStr))
	return api.SendPacket(conn, response)
}

// UpperCaseString makes a string upper case
func (s *Server) UpperCaseString(req *api.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	response := &api.PacketBody{
		MessageType: req.MessageType,
		Data:        []byte(strings.ToUpper(input)),
	}
	s.log.Println("string upper case", input, "->", strings.ToUpper(input))
	return api.SendPacket(conn, response)
}

// LowerCaseString makes a string lower case
func (s *Server) LowerCaseString(req *api.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	response := &api.PacketBody{
		MessageType: req.MessageType,
		Data:        []byte(strings.ToLower(input)),
	}
	s.log.Println("string lower case", input, "->", strings.ToLower(input))
	return api.SendPacket(conn, response)
}

// CountString makes a string lower case
func (s *Server) CountString(req *api.PacketBody, conn net.Conn) error {
	input := string(req.Data)
	words := strings.Split(input, " ")
	letters := 0
	for _, w := range words {
		letters += len(w)
	}
	result := api.CountString{
		Words:   len(words),
		Letters: letters,
	}
	resByte, err := json.Marshal(result)
	if err != nil {
		return err
	}
	response := &api.PacketBody{
		MessageType: req.MessageType,
		Data:        resByte,
	}
	s.log.Println("count string", input, fmt.Sprintf("words:%d , letters:%d", result.Words, result.Letters))
	return api.SendPacket(conn, response)
}
