package socks5

import (
	"context"
	"fmt"
	"net"
	"time"

	netpkg "github.com/interplanetary-network/proxyd/internal/net"

	"github.com/go-gost/gosocks5"
	log "github.com/sirupsen/logrus"
)

func (h *socks5Handler) handleConnect(ctx context.Context, conn net.Conn, network, address string) error {
	flog := log.WithFields(log.Fields{
		"dst": fmt.Sprintf("%s/%s", address, network),
		"cmd": "connect",
	})

	flog.Debugf("%s >> %s", conn.RemoteAddr(), address)

	if h.options.Bypass != nil && h.options.Bypass.Contains(address) {
		resp := gosocks5.NewReply(gosocks5.NotAllowed, nil)
		flog.Trace(resp)
		flog.Debug("bypass: ", address)
		return resp.Write(conn)
	}

	var d net.Dialer
	cc, err := d.DialContext(ctx, network, address)
	if err != nil {
		resp := gosocks5.NewReply(gosocks5.NetUnreachable, nil)
		flog.Trace(resp)
		resp.Write(conn)
		return err
	}

	defer cc.Close()

	resp := gosocks5.NewReply(gosocks5.Succeeded, nil)
	flog.Trace(resp)
	if err := resp.Write(conn); err != nil {
		flog.Error(err)
		return err
	}

	t := time.Now()
	flog.Debugf("%s <-> %s", conn.RemoteAddr(), address)
	netpkg.Transport(conn, cc)
	flog.WithFields(map[string]any{
		"duration": time.Since(t),
	}).Debugf("%s >-< %s", conn.RemoteAddr(), address)

	return nil
}
