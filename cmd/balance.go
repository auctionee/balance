package main

import (
	"github.com/auctionee/balance/pkg/logger"
	"github.com/auctionee/balance/pkg/server"
	"sync"
)

const DEFAULT_PORT = 8080

func main() {
	wg := sync.WaitGroup{}
	s := server.NewAuthServer(DEFAULT_PORT)
	wg.Add(1)
	go func() {
		logger.GetLogger(s.Ctx).Println("Listening on:", DEFAULT_PORT)
		s.Start()
	}()
	wg.Wait()
}
