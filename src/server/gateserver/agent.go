package gateserver

import (
	"fmt"
	"net"
	"server/global"
)

type GateAgent struct {
	conn     net.Conn
	recvChan chan *global.Packet
	sendChan chan *global.Packet

	CloseChan chan bool
	isClosed  bool
}

func createGateAgent(conn net.Conn) *GateAgent {
	agent := new(GateAgent)
	agent.conn = conn
	agent.isClosed = false
	agent.recvChan = make(chan *global.Packet)
	agent.sendChan = make(chan *global.Packet)
	return agent
}

func (self *GateAgent) RecvPacket() {
	buf := make([]byte, 1024)
	conn := self.conn
	for {
		lenght, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error: read failed. " + err.Error())
			self.Close()
			return
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		fmt.Println("Rec[", conn.RemoteAddr().String(), "] Say :", string(buf[0:lenght]))
		reciveStr := string(buf[0:lenght])
		packet := new(global.Packet)
		packet.Context = reciveStr
		self.recvChan <- packet
	}
}

func (self *GateAgent) Close() {
	if self.isClosed {
		return
	}
	self.isClosed = true
	self.CloseChan <- true
}
func (self *GateAgent) SendPacket() {
	for {
		select {
		case packet := <-self.sendChan:
			{
				fmt.Println("send packet: ", packet.Context)
				self.conn.Write([]byte(packet.Context))
			}

		}
	}
}
func (self *GateAgent) HandlePacket() {
	for {
		select {
		case packet := <-self.recvChan:
			{
				fmt.Println("recv packet: ", packet.Context)
				self.sendChan <- packet
			}

		}
	}
}
