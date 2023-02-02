package client

import (
	"encoding/json"
	"fmt"
	"github.com/Gictorbit/textsocket/api"
)

func (c *Client) SendTextReceiveResult(msgType api.MessageType) error {
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
	fmt.Println("result:\n", string(response.Data))
	return nil
}

func (c *Client) CountString(msgType api.MessageType) error {
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
	result := api.CountString{}
	if e := json.Unmarshal(response.Data, &result); e != nil {
		return e
	}
	fmt.Printf("Words:%d\nletters:%d\n", result.Words, result.Letters)
	return nil
}
