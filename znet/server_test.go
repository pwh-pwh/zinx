package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error: ", err)
			return
		}
		fmt.Printf("server call back : %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("[zinx v0.1]")
	go ClientTest()
	s.Serve()
}
