package socks5

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"

	"github.com/go-gost/gosocks5"
	"github.com/interplanetary-network/proxyd/internal/net/udp"
	log "github.com/sirupsen/logrus"
)

func (h *socks5Handler) handleUDP(ctx context.Context, conn net.Conn, address string) error {
	flog := log.WithFields(map[string]any{
		"cmd": "udp",
	})

	if !h.options.EnableUDP {
		reply := gosocks5.NewReply(gosocks5.NotAllowed, nil)
		flog.Trace(reply)
		flog.Error("socks5: UDP relay is disabled")
		return reply.Write(conn)
	}

	cc, err := net.ListenUDP("udp", nil)
	if err != nil {
		flog.Error(err)
		reply := gosocks5.NewReply(gosocks5.Failure, nil)
		flog.Trace(reply)
		reply.Write(conn)
		return err
	}
	defer cc.Close()

	saddr := gosocks5.Addr{}
	saddr.ParseFrom(cc.LocalAddr().String())
	saddr.Type = 0
	saddr.Host, _, _ = net.SplitHostPort(conn.LocalAddr().String()) // replace the IP to the out-going interface's
	reply := gosocks5.NewReply(gosocks5.Succeeded, &saddr)
	log.Trace(reply)
	if err := reply.Write(conn); err != nil {
		log.Error(err)
		return err
	}

	flog1 := flog.WithFields(map[string]any{
		"bind": fmt.Sprintf("%s/%s", cc.LocalAddr(), cc.LocalAddr().Network()),
	})
	flog1.Debugf("bind on %s OK", cc.LocalAddr())

	// obtain a udp connection
	c, err := net.ListenUDP("udp", nil) // UDP association
	if err != nil {
		flog1.Error(err)
		return err
	}
	defer c.Close()

	r := udp.NewRelay(UDPConn(cc, h.options.UDPBufferSize), c).
		WithBypass(h.options.Bypass).
		WithLogger(flog1)
	r.SetBufferSize(h.options.UDPBufferSize)

	go r.Run()

	t := time.Now()
	flog1.Debugf("%s <-> %s", conn.RemoteAddr(), cc.LocalAddr())
	io.Copy(ioutil.Discard, conn)
	flog1.WithFields(map[string]any{"duration": time.Since(t)}).
		Debugf("%s >-< %s", conn.RemoteAddr(), cc.LocalAddr())

	return nil
}
