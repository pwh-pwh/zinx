package utils

import (
	"encoding/json"
	"github.com/pwh-pwh/zinx/ziface"
	"os"
)

type GlobalObj struct {
	TcpServer        ziface.IServer
	Host             string
	TcpPort          int
	Name             string
	Version          string
	MaxPacketSize    uint32
	MaxConn          int
	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32
	ConfFilePath     string
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := os.ReadFile(g.ConfFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		ConfFilePath:     "conf/zinx.json",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}

	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
