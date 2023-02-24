package socks5

import (
	"bytes"
	"net"

	"github.com/go-gost/core/common/bufpool"
	"github.com/go-gost/gosocks5"
)

var (
	DefaultBufferSize = 4096
)

type udpConn struct {
	net.PacketConn
	raddr      net.Addr
	taddr      net.Addr
	bufferSize int
}

func UDPConn(c net.PacketConn, bufferSize int) net.PacketConn {
	return &udpConn{
		PacketConn: c,
		bufferSize: bufferSize,
	}
}

// ReadFrom reads an UDP datagram.
// NOTE: for server side,
// the returned addr is the target address the client want to relay to.
func (c *udpConn) ReadFrom(b []byte) (n int, addr net.Addr, err error) {
	rbuf := bufpool.Get(c.bufferSize)
	defer bufpool.Put(rbuf)

	n, c.raddr, err = c.PacketConn.ReadFrom(*rbuf)
	if err != nil {
		return
	}

	socksAddr := gosocks5.Addr{}
	header := gosocks5.UDPHeader{
		Addr: &socksAddr,
	}
	hlen, err := header.ReadFrom(bytes.NewReader((*rbuf)[:n]))
	if err != nil {
		return
	}
	n = copy(b, (*rbuf)[hlen:n])

	addr, err = net.ResolveUDPAddr("udp", socksAddr.String())
	return
}

func (c *udpConn) Read(b []byte) (n int, err error) {
	n, _, err = c.ReadFrom(b)
	return
}

func (c *udpConn) WriteTo(b []byte, addr net.Addr) (n int, err error) {
	wbuf := bufpool.Get(c.bufferSize)
	defer bufpool.Put(wbuf)

	socksAddr := gosocks5.Addr{}
	if err = socksAddr.ParseFrom(addr.String()); err != nil {
		return
	}

	header := gosocks5.UDPHeader{
		Addr: &socksAddr,
	}
	dgram := gosocks5.UDPDatagram{
		Header: &header,
		Data:   b,
	}

	buf := bytes.NewBuffer((*wbuf)[:0])
	_, err = dgram.WriteTo(buf)
	if err != nil {
		return
	}

	_, err = c.PacketConn.WriteTo(buf.Bytes(), c.raddr)
	n = len(b)

	return
}

func (c *udpConn) Write(b []byte) (n int, err error) {
	return c.WriteTo(b, c.taddr)
}

func (c *udpConn) RemoteAddr() net.Addr {
	return c.raddr
}
