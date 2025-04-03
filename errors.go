package disgomsg

import "errors"

var (
	ErrMissingChannelID   = errors.New("missing channel ID")
	ErrMissingMessageID   = errors.New("missing message ID")
	ErrInvalidMessageType = errors.New("invalid message type")
)
