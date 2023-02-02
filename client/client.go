package client

import (
	"fmt"
	"github.com/Gictorbit/textsocket/api"
	"log"
	"net"
	"sync"
)

type Client struct {
	listenAddr string
	conn       net.Conn
	wg         sync.WaitGroup
	log        *log.Logger
}
type CliInterFace interface {
	Connect() error
	Stop()
}

var (
	_ CliInterFace = &Client{}
)

func NewClient(listenAddr string, logger *log.Logger) CliInterFace {
	return &Client{
		listenAddr: listenAddr,
		wg:         sync.WaitGroup{},
		log:        logger,
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial(api.SocketType, c.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to dial server: %v\n", err.Error())
	}
	c.conn = conn
	return nil
}

func (c *Client) Stop() {
	c.wg.Wait()
	c.conn.Close()
	c.log.Println("stop client...")
}
