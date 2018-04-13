package main

import (
	"fmt"
	//"server/global"
	//"server/config"
	"server/loginserver"
)

//var isQuit = global.GetQuitChannel()

func main() {
	fmt.Println("app start.")
	go loginserver.StartListen()
	<-make(chan int)
}
