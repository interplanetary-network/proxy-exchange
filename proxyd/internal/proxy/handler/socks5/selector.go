package socks5

import (
	"net"

	"github.com/interplanetary-network/proxyd/internal/auth"

	"github.com/go-gost/gosocks5"
	log "github.com/sirupsen/logrus"
)

type serverSelector struct {
	methods       []uint8
	authenticator auth.Authenticator
}

func NewServerSelector(authenticator auth.Authenticator) gosocks5.Selector {
	return &serverSelector{
		authenticator: authenticator,
	}
}

func (selector *serverSelector) Methods() []uint8 {
	return selector.methods
}

func (s *serverSelector) Select(methods ...uint8) (method uint8) {
	log.Debugf("%d %d %v", gosocks5.Ver5, len(methods), methods)
	method = gosocks5.MethodNoAuth

	// when Authenticator is set, auth is mandatory
	if s.authenticator != nil {
		if method == gosocks5.MethodNoAuth {
			method = gosocks5.MethodUserPass
		}
	}

	return
}

func (s *serverSelector) OnSelected(method uint8, conn net.Conn) (net.Conn, error) {
	log.Debugf("%d %d", gosocks5.Ver5, method)
	switch method {
	case gosocks5.MethodUserPass:
		req, err := gosocks5.ReadUserPassRequest(conn)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		log.Trace(req)

		if err := s.authenticator.Authenticate(req.Username, req.Password); err != nil {
			resp := gosocks5.NewUserPassResponse(gosocks5.UserPassVer, gosocks5.Failure)
			if err := resp.Write(conn); err != nil {
				log.Error(err)
				return nil, err
			}
			log.Error(err)

			return nil, gosocks5.ErrAuthFailure
		}

		resp := gosocks5.NewUserPassResponse(gosocks5.UserPassVer, gosocks5.Succeeded)
		log.Trace(resp)
		if err := resp.Write(conn); err != nil {
			log.Error(err)
			return nil, err
		}

	case gosocks5.MethodNoAcceptable:
		return nil, gosocks5.ErrBadMethod
	}

	return conn, nil
}
