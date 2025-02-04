package a2sqc_types

import "net"

type Request struct {
	Data []byte
	Addr net.Addr
}
