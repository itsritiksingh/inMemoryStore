package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/itsritiksingh/inMemoryStore/cmd"
	"github.com/itsritiksingh/inMemoryStore/pkg/sstable"
	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

func main() {
	store := store.Init()
	
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "store", store))

	defer cancel()
	cmd.Execute(ctx)

	systemExitChan := make(chan os.Signal, 1)
	signal.Notify(systemExitChan,syscall.SIGINT,syscall.SIGTERM)

	go sstable.InitTable(store)

	<-systemExitChan
	
}