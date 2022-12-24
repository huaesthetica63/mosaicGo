package main

import (
	"main/mosaic_server"
	"os"
	"os/signal"
)

func main() {
	server := mosaic_server.Server{}
	go server.Load()
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt)
	<-finish
}
