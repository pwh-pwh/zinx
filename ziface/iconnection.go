package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConnID() uint32
	GetTCPConnection() *net.TCPConn
	RemoteAddr() net.Addr
	SendMsg(msgId uint32, data []byte) error
	SendBuffMsg(msgId uint32, data []byte) error
}
