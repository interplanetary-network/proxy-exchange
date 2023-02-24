package socks5

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-gost/gosocks5"
	netpkg "github.com/interplanetary-network/proxyd/internal/net"
	log "github.com/sirupsen/logrus"
)

func (h *socks5Handler) handleBind(ctx context.Context, conn net.Conn, network, address string) error {
	logger := log.WithFields(log.Fields{
		"dst": fmt.Sprintf("%s/%s", address, network),
		"cmd": "bind",
	})

	logger.Debugf("%s >> %s", conn.RemoteAddr(), address)

	if !h.options.EnableBind {
		reply := gosocks5.NewReply(gosocks5.NotAllowed, nil)
		logger.Trace(reply)
		log.Error("socks5: BIND is disabled")
		return reply.Write(conn)
	}

	// BIND does not support chain.
	return h.bindLocal(ctx, conn, network, address)
}

func (h *socks5Handler) bindLocal(ctx context.Context, conn net.Conn, network, address string) error {
	ln, err := net.Listen(network, address) // strict mode: if the port already in use, it will return error
	if err != nil {
		log.Error(err)
		reply := gosocks5.NewReply(gosocks5.Failure, nil)
		if err := reply.Write(conn); err != nil {
			log.Error(err)
		}
		log.Debug(reply)
		return err
	}

	socksAddr := gosocks5.Addr{}
	if err := socksAddr.ParseFrom(ln.Addr().String()); err != nil {
		log.Warn(err)
	}

	// Issue: may not reachable when host has multi-interface
	socksAddr.Host, _, _ = net.SplitHostPort(conn.LocalAddr().String())
	socksAddr.Type = 0
	reply := gosocks5.NewReply(gosocks5.Succeeded, &socksAddr)
	log.Trace(reply)
	if err := reply.Write(conn); err != nil {
		log.Error(err)
		ln.Close()
		return err
	}

	log.Debugf("bind on %s OK", ln.Addr())

	h.serveBind(ctx, conn, ln)
	return nil
}

func (h *socks5Handler) serveBind(ctx context.Context, conn net.Conn, ln net.Listener) {
	flog := log.WithFields(log.Fields{
		"bind": fmt.Sprintf("%s/%s", ln.Addr(), ln.Addr().Network()),
	})

	var rc net.Conn
	accept := func() <-chan error {
		errc := make(chan error, 1)

		go func() {
			defer close(errc)
			defer ln.Close()

			c, err := ln.Accept()
			if err != nil {
				errc <- err
			}
			rc = c
		}()

		return errc
	}

	pc1, pc2 := net.Pipe()
	pipe := func() <-chan error {
		errc := make(chan error, 1)

		go func() {
			defer close(errc)
			defer pc1.Close()

			errc <- netpkg.Transport(conn, pc1)
		}()

		return errc
	}

	defer pc2.Close()

	select {
	case err := <-accept():
		if err != nil {
			flog.Error(err)

			reply := gosocks5.NewReply(gosocks5.Failure, nil)
			flog.Trace(reply)
			if err := reply.Write(pc2); err != nil {
				flog.Error(err)
			}

			return
		}
		defer rc.Close()

		flog.Debugf("peer %s accepted", rc.RemoteAddr())

		logger1 := flog.WithFields(log.Fields{
			"local":  rc.LocalAddr().String(),
			"remote": rc.RemoteAddr().String(),
		})

		raddr := gosocks5.Addr{}
		raddr.ParseFrom(rc.RemoteAddr().String())
		reply := gosocks5.NewReply(gosocks5.Succeeded, &raddr)
		logger1.Trace(reply)
		if err := reply.Write(pc2); err != nil {
			logger1.Error(err)
		}

		start := time.Now()
		logger1.Debugf("%s <-> %s", rc.LocalAddr(), rc.RemoteAddr())
		netpkg.Transport(pc2, rc)
		logger1.WithFields(map[string]any{"duration": time.Since(start)}).
			Debugf("%s >-< %s", rc.LocalAddr(), rc.RemoteAddr())

	case err := <-pipe():
		if err != nil {
			flog.Error(err)
		}
		ln.Close()
		return
	}
}
