//go:build with_utls

package vless

import (
	"net"
	"reflect"
	"unsafe"

	N "github.com/sagernet/sing/common/network"

	ec "github.com/domaingts/electricity"
	venc "github.com/domaingts/venc"
	utls "github.com/metacubex/utls"
)

func init() {
	tlsRegistry = append(tlsRegistry, func(conn net.Conn) (loaded bool, netConn net.Conn, reflectType reflect.Type, reflectPointer uintptr) {
		uConn, loaded := N.CastReader[*ec.Conn](conn)
		if loaded {
			return true, uConn.NetConn(), reflect.TypeOf(uConn).Elem(), uintptr(unsafe.Pointer(uConn))
		}
		decryptionConn, loaded := N.CastReader[*venc.CommonConn](conn)
		if loaded {
			return true, decryptionConn.Conn, reflect.TypeOf(decryptionConn).Elem(), uintptr(unsafe.Pointer(decryptionConn))
		}
		tlsConn, loaded := N.CastReader[*utls.Conn](conn)
		if loaded {
			return true, tlsConn.NetConn(), reflect.TypeOf(tlsConn).Elem(), uintptr(unsafe.Pointer(tlsConn))
		}
		return
	})
}
