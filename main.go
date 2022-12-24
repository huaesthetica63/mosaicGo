package main

import (
	"main/mosaic_server"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := mosaic_server.NewServer(":8080", time.Second*5)
	go server.Load()
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt)
	<-finish
}
