package handler

import (
	"context"
	"math"
	"net"
	"time"

	"github.com/interplanetary-network/proxyd/internal/auth"
	"github.com/interplanetary-network/proxyd/internal/bypass"
)

type Options struct {
	Bypass        bypass.Bypass
	ReadTimeout   time.Duration
	Authenticator auth.Authenticator
	EnableUDP     bool
	EnableBind    bool
	UDPBufferSize int
}

type Option func(*Options)

func WithBypass(b bypass.Bypass) Option {
	return func(o *Options) {
		o.Bypass = b
	}
}

func WithEnableUDP(enable bool) Option {
	return func(o *Options) {
		o.EnableUDP = enable
	}
}

func WithEnableBind(enable bool) Option {
	return func(o *Options) {
		o.EnableBind = enable
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = timeout
	}
}

func WithUDPBufferSize(size int) Option {
	return func(o *Options) {
		o.UDPBufferSize = int(math.Min(math.Max(float64(size), 512), 64*1024))
	}
}

func WithAuthenticator(authenticator auth.Authenticator) Option {
	return func(o *Options) {
		o.Authenticator = authenticator
	}
}

type Handler interface {
	Handle(ctx context.Context, conn net.Conn) error
}
