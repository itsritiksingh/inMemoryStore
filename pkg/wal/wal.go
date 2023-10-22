// write ahead log
package wal

import (
	"log"
	"os"
)

type Wal struct {
	f *os.File
}

func Init() (*Wal)  {
	fileName := "wal"
	_ , error := os.Stat(fileName)

	// check if error is "file not exists"
	if !os.IsNotExist(error) {
		if err := os.Remove("testfile.txt"); err != nil {
			log.Fatal(err)
		}
	}

	file , err := createWrite(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return &Wal{
		f: file,
	}
}

func createWrite(fileName string) (*os.File,error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		
	return f,err
}

func (w *Wal) Write(write []byte){
	if _, err := w.f.Write(write); err != nil {
		w.f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := w.f.Close(); err != nil {
		log.Fatal(err)
	}
}