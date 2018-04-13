package main

import (
	"fmt"
	//"server/global"
	//"server/config"
	"server/gateserver"
	"server/loginserver"
)

//var isQuit = global.GetQuitChannel()

func main() {
	fmt.Println("app start.")
	go loginserver.StartListen()
	go gateserver.StartListen()
	<-make(chan int)
}
