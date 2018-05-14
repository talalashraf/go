package mydb

import (
	"errors"
	"fmt"
)

var (
	// ErrClientNotInitialized says client hasn't been initialized
	ErrClientNotInitialized = errors.New("client has not been initialized")
)

// Client is a dummy client
type Client struct {
	Conn string
}

// Do implements a client do
func (c *Client) Do(cmd string) error {
	if c.Conn == "" {
		return ErrClientNotInitialized
	}
	fmt.Printf("Client %s -> %s\n", c.Conn, cmd)
	return nil
}

// Init initializes
func Init(s string) *Client {
	return &Client{Conn: s}
}
