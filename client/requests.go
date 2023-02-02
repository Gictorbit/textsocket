package client

import (
	"fmt"
	"github.com/Gictorbit/textsocket/api"
)

func (c *Client) ReverseString(msgType api.MessageType) error {
	textInput, err := c.ReadPrompt()
	if err != nil {
		return err
	}
	request := &api.PacketBody{
		MessageType: msgType,
		Data:        []byte(textInput),
	}
	if e := api.SendPacket(c.conn, request); e != nil {
		return e
	}
	response, err := api.ReadPacket(c.conn)
	if err != nil {
		return err
	}
	fmt.Println("result ->", string(response.Data))
	return nil
}
