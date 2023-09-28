package ip

import (
	"net"
	"net/http"
	"strings"
)

var ip string

var r http.Request

var localnetWorks = []string{
	"10.0.0.0/8",
	"169.254.0.0/16",
	"172.16.0.0/12",
	"172.17.0.0/12",
	"172.18.0.0/12",
	"172.19.0.0/12",
	"172.20.0.0/12",
	"172.21.0.0/12",
	"172.22.0.0/12",
	"172.23.0.0/12",
	"172.24.0.0/12",
	"172.25.0.0/12",
	"172.26.0.0/12",
	"172.27.0.0/12",
	"172.28.0.0/12",
	"172.29.0.0/12",
	"172.30.0.0/12",
	"172.31.0.0/12",
	"192.168.0.0/16",
}

func ClientIP() string {
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !HasLocalIPAddr(ip) {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIPAddr(ip) {
		return ip
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		if !HasLocalIPAddr(ip) {
			return ip
		}
	}

	return ""
}

func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

func HasLocalIP(ip net.IP) bool {
	for _, netWork := range localnetWorks {
		if strings.Contains(netWork, ip.String()) {
			return true
		}
	}

	return ip.IsLoopback()
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
