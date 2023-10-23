package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/itsritiksingh/inMemoryStore/pkg/sstable"
	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

func main() {
	db := store.Init()

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	systemExitChan := make(chan os.Signal, 1)
	signal.Notify(systemExitChan,syscall.SIGINT,syscall.SIGTERM)

	go sstable.InitTable(db)

	_,_ = db.Put("hello","world")
	fmt.Println(db.Get("hello"))

	<-systemExitChan
	
}