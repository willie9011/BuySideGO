package AppInit

import (
	"log"
	"os"
	"os/signal"
)

var ServerSigChannel chan os.Signal

func init() {
	ServerSigChannel = make(chan os.Signal)
}

func ServerNotify() {
	signal.Notify(ServerSigChannel, os.Interrupt)
	<-ServerSigChannel
}

func ShutDownServer(err error) {
	log.Println(err)
	ServerSigChannel <- os.Interrupt
}
