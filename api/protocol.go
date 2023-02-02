package api

import (
	"errors"
	"io"
	"net"
)

const (
	SocketType          = "tcp"
	PacketMaxByteLength = 2048
	HostAddress         = "127.0.0.1"
	PortNumber          = 2023
)

var (
	ErrInvalidPacketSize = errors.New("invalid packet size")
)

type MessageType uint8

const (
	MessageTypeReverseString   MessageType = 1
	MessageTypeUpperCaseString MessageType = 2
	MessageTypeLowerCaseString MessageType = 3
	MessageTypeCountString     MessageType = 4
)

type PacketBody struct {
	MessageType MessageType
	Data        []byte
}

func ReadPacket(conn net.Conn) (*PacketBody, error) {
	buf := make([]byte, PacketMaxByteLength)
	n, err := conn.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	return &PacketBody{
		MessageType: MessageType(buf[0]),
		Data:        buf[1:n],
	}, nil
}

func SendPacket(conn net.Conn, packet *PacketBody) error {
	buf := make([]byte, 0)
	buf = append(buf, byte(packet.MessageType))
	buf = append(buf, packet.Data...)
	if len(buf) > PacketMaxByteLength {
		return ErrInvalidPacketSize
	}
	if _, err := conn.Write(buf); err != nil {
		return err
	}
	return nil
}

type CountString struct {
	Letters int `json:"letters"`
	Words   int `json:"words"`
}
