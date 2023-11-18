// write ahead log
package wal

import (
	"log"
	"os"
)

const (
	WalFileName = "wal"
)

func Write(write []byte){
	f, err := os.OpenFile(WalFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		
	if _, err := f.Write(write); err != nil {
		defer f.Close() 
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}