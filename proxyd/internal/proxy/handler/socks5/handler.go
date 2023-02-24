package socks5

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/interplanetary-network/proxyd/internal/proxy/handler"

	"github.com/go-gost/gosocks5"
	log "github.com/sirupsen/logrus"
)

var (
	ErrUnknownCmd = errors.New("socks5: unknown command")
)

type socks5Handler struct {
	selector gosocks5.Selector
	options  *handler.Options
}

func NewHandler(selector gosocks5.Selector, opts ...handler.Option) handler.Handler {
	options := &handler.Options{
		// Default 4096
		UDPBufferSize: 4096,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &socks5Handler{
		selector: selector,
		options:  options,
	}
}

func (h *socks5Handler) Handle(ctx context.Context, conn net.Conn) error {
	defer conn.Close()

	start := time.Now()

	flog := log.WithFields(map[string]any{
		"remote": conn.RemoteAddr().String(),
		"local":  conn.LocalAddr().String(),
	})

	flog.Infof("%s <> %s", conn.RemoteAddr(), conn.LocalAddr())
	defer func() {
		flog.WithFields(map[string]any{
			"duration": time.Since(start),
		}).Infof("%s >< %s", conn.RemoteAddr(), conn.LocalAddr())
	}()

	if h.options.ReadTimeout > 0 {
		conn.SetReadDeadline(time.Now().Add(h.options.ReadTimeout))
	}

	conn = gosocks5.ServerConn(conn, h.selector)
	req, err := gosocks5.ReadRequest(conn)
	if err != nil {
		flog.Error(err)
		return err
	}
	flog.Trace(req)
	conn.SetReadDeadline(time.Time{})

	address := req.Addr.String()

	switch req.Cmd {
	case gosocks5.CmdConnect:
		return h.handleConnect(ctx, conn, "tcp", address)
	case gosocks5.CmdBind:
		return h.handleBind(ctx, conn, "tcp", address)
	case gosocks5.CmdUdp:
		return h.handleUDP(ctx, conn, address)
	default:
		err = ErrUnknownCmd
		flog.Error(err)
		resp := gosocks5.NewReply(gosocks5.CmdUnsupported, nil)
		flog.Trace(resp)
		resp.Write(conn)
		return err
	}
}
