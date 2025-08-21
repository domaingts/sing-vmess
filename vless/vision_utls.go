//go:build with_utls

package vless

import (
	"net"
	"reflect"
	"unsafe"

	"github.com/sagernet/sing/common"

	utls "github.com/metacubex/utls"

	ec "github.com/domaingts/electricity"
)

func init() {
	tlsRegistry = append(tlsRegistry, func(conn net.Conn) (loaded bool, netConn net.Conn, reflectType reflect.Type, reflectPointer uintptr) {
		tlsConn, loaded := common.Cast[*ec.Conn](conn)
		if loaded {
			return true, tlsConn.NetConn(), reflect.TypeOf(tlsConn).Elem(), uintptr(unsafe.Pointer(tlsConn))
		}
		uConn, loaded := common.Cast[*utls.UConn](conn)
		if loaded {
			return true, uConn.NetConn(), reflect.TypeOf(uConn.Conn).Elem(), uintptr(unsafe.Pointer(uConn.Conn))
		}
		return
	})
}
