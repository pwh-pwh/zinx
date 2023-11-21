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
	//设置链接属性
	SetProperty(key string, value any)
	//获取链接属性
	GetProperty(key string) (any, error)
	//移除链接属性
	RemoveProperty(key string)
}
