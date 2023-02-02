package client

import (
	"bufio"
	"fmt"
	"github.com/Gictorbit/textsocket/api"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Client struct {
	listenAddr string
	conn       net.Conn
	wg         sync.WaitGroup
	log        *log.Logger
	stdin      *bufio.Reader
}
type CliInterFace interface {
	Connect() error
	ReadOption() (api.MessageType, error)
	Stop()
	SendTextReceiveResult(msgType api.MessageType) error
	CountString(msgType api.MessageType) error
}

var (
	_ CliInterFace = &Client{}
)

func NewClient(listenAddr string, logger *log.Logger) CliInterFace {
	reader := bufio.NewReader(os.Stdin)
	return &Client{
		listenAddr: listenAddr,
		wg:         sync.WaitGroup{},
		log:        logger,
		stdin:      reader,
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

func (c *Client) ReadOption() (api.MessageType, error) {
	c.PrintOptions()
	fmt.Printf("-> ")
	text, err := c.stdin.ReadString('\n')
	if err != nil {
		return 0, err
	}
	text = strings.Replace(text, "\n", "", -1)
	msgType, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return api.MessageType(msgType), nil
}

func (c *Client) ReadPrompt() (string, error) {
	fmt.Printf("-> ")
	text, err := c.stdin.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)
	return text, nil
}

func (c *Client) PrintOptions() {
	fmt.Println("")
	fmt.Println("===============| Choose an option |===============")
	for name, msgType := range api.MessageTypes {
		fmt.Printf("%d- %s\n", msgType, name)
	}
	fmt.Println("==================================================")
}
