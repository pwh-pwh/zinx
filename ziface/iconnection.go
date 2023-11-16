package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConnID() uint32
	GetTCPConnection() *net.TCPConn
}

type HandFunc func(*net.TCPConn, []byte, int) error
