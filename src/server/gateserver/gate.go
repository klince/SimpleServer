package gateserver

import (
	"fmt"
	"net"
)

func StartListen() {
	fmt.Println("gate server start.")
	service := ":9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println("error: gate server ResolveTCPAddr failed.", service)
		return
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("error: gate server ListenTCP failed.", tcpAddr.String())
		return
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("error: gate server Accept failed.")
		return
	}
	go Handler(conn) //此处使用go关键字新建线程处理连接，实现并发
}
func Handler(conn net.Conn) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())
	defer conn.Close()
	agent := createGateAgent(conn)
	go agent.RecvPacket()
	go agent.SendPacket()
	go agent.HandlePacket()
	<-agent.CloseChan
}
