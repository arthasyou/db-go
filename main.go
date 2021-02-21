package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/arthasyou/db-go/mysql"
)

func main() {
	// mongo.Connect("localhost", 27017, "", "")
	mysql.Connect("localhost", 3306, "root", "123456", "test")
	// mysql.Query()
	waitExit()
}

func waitExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
}
