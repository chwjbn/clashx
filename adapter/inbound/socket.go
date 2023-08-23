package inbound

import (
	"net"

	C "github.com/chwjbn/clashx/constant"
	"github.com/chwjbn/clashx/context"
	"github.com/chwjbn/clashx/transport/socks5"
)

// NewSocket receive TCP inbound and return ConnContext
func NewSocket(target socks5.Addr, conn net.Conn, source C.Type) *context.ConnContext {
	metadata := parseSocksAddr(target)
	metadata.NetWork = C.TCP
	metadata.Type = source
	if ip, port, err := parseAddr(conn.RemoteAddr().String()); err == nil {
		metadata.SrcIP = ip
		metadata.SrcPort = port
	}

	return context.NewConnContext(conn, metadata)
}


func NewSocketWithAuth(target socks5.Addr,proxy_user string,proxy_pass string, conn net.Conn, source C.Type) *context.ConnContext {
	metadata := parseSocksAddr(target)
	metadata.NetWork = C.TCP
	metadata.Type = source
	metadata.ProxyUser=proxy_user
	metadata.ProxyPass=proxy_pass
	if ip, port, err := parseAddr(conn.RemoteAddr().String()); err == nil {
		metadata.SrcIP = ip
		metadata.SrcPort = port
	}

	return context.NewConnContext(conn, metadata)
}