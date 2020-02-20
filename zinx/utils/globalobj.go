package utils

import (
	"encoding/json"
	"io/ioutil"
	// "os"
	"zinx/ziface"
	// "zinx/zlog"
)

// storage zinx config

type GlobalObj struct {
	// server
	TcpServer ziface.IServer // current server obj
	Host      string         // current server host
	TcpPort   int            //
	Name      string         //

	// zinx
	Version        string // zinx version
	MaxConn        int    //
	MaxPackageSize uint32 //
}

// define global obj
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// extract json to struct
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// init global obj
func init() {
	// default
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// read conf/zinx.json
	GlobalObject.Reload()
}
