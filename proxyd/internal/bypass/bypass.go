package bypass

import (
	"net"
	"strings"

	"github.com/interplanetary-network/proxyd/internal/matcher"
)

type Bypass interface {
	Contains(addr string) bool
}

type bypass struct {
	ip       matcher.Matcher
	cidr     matcher.Matcher
	domain   matcher.Matcher
	wildcard matcher.Matcher
}

func NewBypass(pattern ...string) Bypass {
	var ips []net.IP
	var inets []*net.IPNet
	var domains []string
	var wildcards []string

	for _, v := range pattern {
		if ip := net.ParseIP(v); ip != nil {
			ips = append(ips, ip)
			continue
		}
		if _, inet, err := net.ParseCIDR(v); err == nil {
			inets = append(inets, inet)
			continue
		}
		if strings.ContainsAny(v, "*?") {
			wildcards = append(wildcards, v)
			continue
		}
		domains = append(domains, v)
	}

	return &bypass{
		ip:       matcher.IPMatcher(ips),
		cidr:     matcher.CIDRMatcher(inets),
		domain:   matcher.DomainMatcher(domains),
		wildcard: matcher.WildcardMatcher(wildcards),
	}
}

func (b *bypass) Contains(addr string) bool {
	if addr == "" {
		return false
	}

	if host, _, _ := net.SplitHostPort(addr); host != "" {
		addr = host
	}

	if ip := net.ParseIP(addr); ip != nil {
		return b.ip.Match(addr) || b.cidr.Match(addr)
	}

	return b.domain.Match(addr) || b.wildcard.Match(addr)
}
