package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/arthasyou/db-go/redis"
)

func main() {
	redis.Connect("localhost:6379", "", 1)
	redis.Set("idihig", "hggyyggig")
	r, _ := redis.Get("idihig")
	fmt.Println(r)
	waitExit()
}

func waitExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
}
