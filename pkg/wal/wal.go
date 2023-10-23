// write ahead log
package wal

import (
	"log"
	"os"
)

type Wal struct {
	f *os.File
}

const (
	walFileName = "wal"
)

func Init() (*Wal)  {
	_ , error := os.Stat(walFileName)

	// check if error is "file not exists"
	if !os.IsNotExist(error) {
		if err := os.Remove(walFileName); err != nil {
			log.Fatal(err)
		}
	}

	file , err := createWrite(walFileName)

	if err != nil {
		log.Fatal(err)
	}

	return &Wal{
		f: file,
	}
}

func createWrite(walFileName string) (*os.File,error) {
	f, err := os.OpenFile(walFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		
	return f,err
}

func (w *Wal) Write(write []byte){
	if _, err := w.f.Write(write); err != nil {
		w.f.Close() 
		log.Fatal(err)
	}
	if err := w.f.Close(); err != nil {
		log.Fatal(err)
	}
}