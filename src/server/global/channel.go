package global

var isQuit chan bool

func GetQuitChannel() chan bool {
	return isQuit
}
