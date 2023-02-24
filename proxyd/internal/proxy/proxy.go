package proxy

import (
	"context"
	"errors"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/interplanetary-network/proxyd/internal/auth"
	"github.com/interplanetary-network/proxyd/internal/bypass"
	"github.com/interplanetary-network/proxyd/internal/proxy/handler"
	"github.com/interplanetary-network/proxyd/internal/proxy/handler/http"
	"github.com/interplanetary-network/proxyd/internal/proxy/handler/socks5"
	log "github.com/sirupsen/logrus"
)

type options struct {
	bypass        []string
	authenticator auth.Authenticator
}

type Option func(*options)

func WithBypass(bypass ...string) Option {
	return func(o *options) {
		o.bypass = bypass
	}
}

func WithAuthenticator(authenticator auth.Authenticator) Option {
	return func(o *options) {
		o.authenticator = authenticator
	}
}

var (
	ErrInvalidProtocol = errors.New("invalid protocol")
)

const (
	SOCKS5 = "socks5"
	HTTP   = "http"
)

var supportedProtocols = []string{SOCKS5, HTTP}

type Service interface {
	Run(ctx context.Context) error
}

type service struct {
	url *url.URL

	listener net.Listener
	handler  handler.Handler
}

func ParseFromURL(rawURL string, opts ...Option) (Service, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	opt := &options{}
	for _, o := range opts {
		o(opt)
	}

	var handlerOptions []handler.Option

	// Parse global options
	if opt.bypass != nil {
		handlerOptions = append(handlerOptions, handler.WithBypass(bypass.NewBypass(opt.bypass...)))
	}

	if opt.authenticator != nil {
		handlerOptions = append(handlerOptions, handler.WithAuthenticator(opt.authenticator))
	}

	// Parse URL query parameters
	if u.Query().Get("udp") == "1" {
		handlerOptions = append(handlerOptions, handler.WithEnableUDP(true))
	}

	if u.Query().Get("bind") == "1" {
		handlerOptions = append(handlerOptions, handler.WithEnableBind(true))
	}

	if u.Query().Get("udp-buffer-size") != "" {
		size, err := strconv.Atoi(u.Query().Get("udp-buffer-size"))
		if err != nil {
			return nil, err
		}
		handlerOptions = append(handlerOptions, handler.WithUDPBufferSize(size))
	}

	if u.Query().Get("read-timeout") != "" {
		readTimeout, err := time.ParseDuration(u.Query().Get("read-timeout"))
		if err != nil {
			return nil, err
		}
		handlerOptions = append(handlerOptions, handler.WithReadTimeout(readTimeout))
	}

	var h handler.Handler
	switch u.Scheme {
	case SOCKS5:
		h = socks5.NewHandler(socks5.NewServerSelector(opt.authenticator), handlerOptions...)
	case HTTP:
		h = http.NewHandler(handlerOptions...)
	default:
		return nil, ErrInvalidProtocol
	}

	addr, err := net.ResolveTCPAddr("tcp", u.Host)
	if err != nil {
		return nil, err
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	log.Info("listening on ", ln.Addr().String())

	return &service{
		url:      u,
		handler:  h,
		listener: ln,
	}, nil
}

func (s *service) Run(ctx context.Context) error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Error(err)
			return err
		}

		go s.handler.Handle(ctx, conn)
	}
}
