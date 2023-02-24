package http

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	netpkg "github.com/interplanetary-network/proxyd/internal/net"
	"github.com/interplanetary-network/proxyd/internal/proxy/handler"

	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"
)

type httpHandler struct {
	options *handler.Options
}

func NewHandler(opts ...handler.Option) handler.Handler {
	options := &handler.Options{}
	for _, opt := range opts {
		opt(options)
	}

	return &httpHandler{
		options: options,
	}
}

func (h *httpHandler) Handle(ctx context.Context, conn net.Conn) error {
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

	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Error(err)
		return err
	}
	defer req.Body.Close()

	return h.handleRequest(ctx, conn, req)
}

func (h *httpHandler) handleRequest(ctx context.Context, conn net.Conn, req *http.Request) error {
	if !req.URL.IsAbs() && govalidator.IsDNSName(req.Host) {
		req.URL.Scheme = "http"
	}

	addr := req.Host
	if _, port, _ := net.SplitHostPort(addr); port == "" {
		addr = net.JoinHostPort(addr, "80")
	}

	fields := log.Fields{
		"dst": addr,
	}
	if u, _, _ := h.basicProxyAuth(req.Header.Get("Proxy-Authorization")); u != "" {
		fields["user"] = u
	}
	flog := log.WithFields(fields)

	if log.IsLevelEnabled(log.TraceLevel) {
		dump, _ := httputil.DumpRequest(req, false)
		log.Trace(string(dump))
	}
	flog.Debugf("%s >> %s", conn.RemoteAddr(), addr)

	resp := &http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	if resp.Header == nil {
		resp.Header = http.Header{}
	}

	if h.options.Bypass != nil && h.options.Bypass.Contains(addr) {
		resp.StatusCode = http.StatusForbidden

		if log.IsLevelEnabled(log.TraceLevel) {
			dump, _ := httputil.DumpResponse(resp, false)
			log.Trace(string(dump))
		}
		flog.Debug("bypass: ", addr)

		return resp.Write(conn)
	}

	if !h.authenticate(conn, req, resp) {
		return nil
	}

	if req.Method == "PRI" ||
		(req.Method != http.MethodConnect && req.URL.Scheme != "http") {
		resp.StatusCode = http.StatusBadRequest

		if log.IsLevelEnabled(log.TraceLevel) {
			dump, _ := httputil.DumpResponse(resp, false)
			log.Trace(string(dump))
		}

		return resp.Write(conn)
	}

	req.Header.Del("Proxy-Authorization")

	var d net.Dialer
	cc, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		resp.StatusCode = http.StatusServiceUnavailable

		if log.IsLevelEnabled(log.TraceLevel) {
			dump, _ := httputil.DumpResponse(resp, false)
			log.Trace(string(dump))
		}
		resp.Write(conn)
		return err
	}
	defer cc.Close()

	if req.Method == http.MethodConnect {
		resp.StatusCode = http.StatusOK
		resp.Status = "200 Connection established"

		if log.IsLevelEnabled(log.TraceLevel) {
			dump, _ := httputil.DumpResponse(resp, false)
			log.Trace(string(dump))
		}
		if err = resp.Write(conn); err != nil {
			log.Error(err)
			return err
		}
	} else {
		req.Header.Del("Proxy-Connection")
		if err = req.Write(cc); err != nil {
			log.Error(err)
			return err
		}
	}

	start := time.Now()
	log.Debugf("%s <-> %s", conn.RemoteAddr(), addr)
	netpkg.Transport(conn, cc)
	log.WithFields(map[string]any{
		"duration": time.Since(start),
	}).Debugf("%s >-< %s", conn.RemoteAddr(), addr)

	return nil
}

func (h *httpHandler) basicProxyAuth(proxyAuth string) (username, password string, ok bool) {
	if proxyAuth == "" {
		return
	}

	if !strings.HasPrefix(proxyAuth, "Basic ") {
		return
	}
	c, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(proxyAuth, "Basic "))
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}

	return cs[:s], cs[s+1:], true
}

func (h *httpHandler) authenticate(conn net.Conn, req *http.Request, resp *http.Response) (ok bool) {
	u, p, _ := h.basicProxyAuth(req.Header.Get("Proxy-Authorization"))

	if err := h.options.Authenticator.Authenticate(u, p); err == nil {
		return true
	}

	if resp.Header == nil {
		resp.Header = http.Header{}
	}

	if resp.StatusCode == 0 {
		realm := "proxyd"
		resp.StatusCode = http.StatusProxyAuthRequired
		resp.Header.Add("Proxy-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", realm))
		if strings.ToLower(req.Header.Get("Proxy-Connection")) == "keep-alive" {
			// XXX libcurl will keep sending auth request in same conn
			// which we don't supported yet.
			resp.Header.Set("Connection", "close")
			resp.Header.Set("Proxy-Connection", "close")
		}

		log.Debug("proxy authentication required")
	} else {
		// resp.Header.Set("Server", "nginx/1.20.1")
		// resp.Header.Set("Date", time.Now().Format(http.TimeFormat))
		if resp.StatusCode == http.StatusOK {
			resp.Header.Set("Connection", "keep-alive")
		}
	}

	if log.IsLevelEnabled(log.TraceLevel) {
		dump, _ := httputil.DumpResponse(resp, false)
		log.Trace(string(dump))
	}

	resp.Write(conn)
	return
}
