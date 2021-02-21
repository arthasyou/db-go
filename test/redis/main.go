package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/luobin998877/go_db/redis"
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
