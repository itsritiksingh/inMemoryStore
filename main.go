package main

import (
	"fmt"

	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

func main() {
	db := store.Init()

	_,_ = db.Put("hello","world")
	fmt.Println(db.Get("hello"))
}