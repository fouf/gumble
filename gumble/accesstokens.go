package gumble // import "layeh.com/gumble/gumble"

import (
	"layeh.com/gumble/gumble/MumbleProto"
)

// AccessTokens are additional passwords that can be provided to the server to
// gain access to restricted channels.
type AccessTokens []string

func (at AccessTokens) writeMessage(client *Client) error {
	packet := MumbleProto.Authenticate{
		Tokens: at,
	}
	return client.Conn.WriteProto(&packet)
}
