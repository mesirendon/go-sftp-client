package sshclient

import "fmt"

type (
	Client struct{}

	Options func(*Client) error
)

func (c *Client) Greet(s string) string {
	return fmt.Sprintf("Hello %s", s)
}
