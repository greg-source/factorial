package main

import (
	"context"
	"github.com/greg-source/factorial/internal"
	"log"
	"os"
	"time"
)

const (
	Port = "8989"
)

func main() {
	server := internal.NewServer(Port)
	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal("failed to start http server")
		}
	}()
	systemQuit := make(chan os.Signal, 1)
	<-systemQuit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	server.Shutdown(ctx)
	cancel()
}
