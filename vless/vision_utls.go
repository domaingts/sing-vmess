//go:build with_utls

package vless

import (
	"net"
	"reflect"
	"sync"
	"unsafe"

	N "github.com/sagernet/sing/common/network"

	ec "github.com/domaingts/electricity"
	venc "github.com/domaingts/venc"
	utls "github.com/metacubex/utls"
)

func init() {
	var (
		ecOffsets = sync.OnceValues(func() (uintptr, uintptr) {
			t := reflect.TypeFor[*ec.Conn]().Elem()
			input, _ := t.FieldByName("input")
			rawInput, _ := t.FieldByName("rawInput")
			return input.Offset, rawInput.Offset
		})
		vencOffsets = sync.OnceValues(func() (uintptr, uintptr) {
			t := reflect.TypeFor[*venc.CommonConn]().Elem()
			input, _ := t.FieldByName("input")
			rawInput, _ := t.FieldByName("rawInput")
			return input.Offset, rawInput.Offset
		})
		utlsOffsets = sync.OnceValues(func() (uintptr, uintptr) {
			t := reflect.TypeFor[*utls.Conn]().Elem()
			input, _ := t.FieldByName("input")
			rawInput, _ := t.FieldByName("rawInput")
			return input.Offset, rawInput.Offset
		})
	)
	tlsRegistry = append(tlsRegistry, func(conn net.Conn) (loaded bool, netConn net.Conn, offsets offsets, reflectPointer unsafe.Pointer) {
		uConn, loaded := N.CastReader[*ec.Conn](conn)
		if loaded {
			return true, uConn.NetConn(), ecOffsets, unsafe.Pointer(uConn)
		}
		decryptionConn, loaded := N.CastReader[*venc.CommonConn](conn)
		if loaded {
			return true, decryptionConn.Conn, vencOffsets, unsafe.Pointer(decryptionConn)
		}
		tlsConn, loaded := N.CastReader[*utls.Conn](conn)
		if loaded {
			return true, tlsConn.NetConn(), utlsOffsets, unsafe.Pointer(tlsConn)
		}
		return
	})
}
