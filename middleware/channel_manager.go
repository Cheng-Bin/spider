package middleware

import (
	"github.com/chapin/spider/base"
)

const (
	CHANNEL_MANAGER_STATUS_UNINITIALIZED ChannelManagerStatus = 0
	CHANNEL_MANAGER_STATUS_INITIALIZED   ChannelManagerStatus = 1
	CHANNEL_MANAGER_STATUS_CLOSED        ChannelManagerStatus = 2
)

// ChannelManagerStatus type.
type ChannelManagerStatus uint8

// ChannelManager interface.
type ChannelManager interface {

	// Init channel.
	Init(channelLen uint, reset bool) bool

	// Close channel.
	Close() bool

	// ReqChan request channel
	ReqChan() (chan base.Request, error)

	// RespChan response channel.
	RespChan() (chan base.Response, error)

	// ItemChan item channel.
	ItemChan() (chan base.Item, error)

	// ErrorChan error channel.
	ErrorChan() (chan error, error)

	// ChannelLen
	ChannelLen() uint

	// Status method.
	Status() ChannelManagerStatus

	// Summary method.
	Summary() string
}
