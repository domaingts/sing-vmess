//go:build with_utls

package vless

import (
	"net"
	"reflect"
	"unsafe"

	N "github.com/sagernet/sing/common/network"

	ec "github.com/domaingts/electricity"
	"github.com/domaingts/venc"
)

func init() {
	tlsRegistry = append(tlsRegistry, func(conn net.Conn) (loaded bool, netConn net.Conn, reflectType reflect.Type, reflectPointer uintptr) {
		tlsConn, loaded := N.CastReader[*ec.Conn](conn)
		if loaded {
			return true, tlsConn.NetConn(), reflect.TypeOf(tlsConn).Elem(), uintptr(unsafe.Pointer(tlsConn))
		}
		decryptionConn, loaded := N.CastReader[*venc.CommonConn](conn)
		if loaded {
			return true, decryptionConn.Conn, reflect.TypeOf(decryptionConn).Elem(), uintptr(unsafe.Pointer(decryptionConn))
		}
		return
	})
}
