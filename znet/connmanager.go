package znet

import (
	"fmt"
	"github.com/pwh-pwh/zinx/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func (c ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Println("Add connID = ", conn.GetConnID(), " success!")
	fmt.Println("conn num = ", c.Len())
}

func (c ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())
	fmt.Println("Remove connID = ", conn.GetConnID(), " success!")
	fmt.Println("conn num = ", c.Len())
}

func (c ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("no found connID = %d", connID)
	}
}

func (c ConnManager) Len() int {
	return len(c.connections)
}

func (c ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
	fmt.Println("Clear all conn success!")
	fmt.Println("conn num = ", c.Len())
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
