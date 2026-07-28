package sshclient

type (
	Client struct{}

	Options func(*Client) error
)
