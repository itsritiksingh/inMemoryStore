package sstable

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

var ssfolder = filepath.Join("pkg","sstable","tables")

func InitTable(s *store.Store) {
	tick := time.NewTicker(10 * time.Minute)

	for {
		<-tick.C
		CreateTable(s)
	}
}

func CreateTable(s *store.Store) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	keySlice := make([]string, 0)
	tmp := make(map[string]string)

	for key, value := range s.Store {
		tmp[key] = value
		keySlice = append(keySlice, key)
	}

	sort.Strings(keySlice)

	err := writeSS(tmp, keySlice)

	if err != nil {
		log.Fatal("couldn't write to log")
	}
}

func writeSS(tmp map[string]string, keySlice []string) error {
	os.MkdirAll(ssfolder,os.ModePerm)
	fileName := fmt.Sprintf("%s/ss log %v_%v_%v",ssfolder, time.Now().Local().Hour(), time.Now().Local().Minute(), time.Now().Local().Second())
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, key := range keySlice {
		value := tmp[key]

		f.Write([]byte(fmt.Sprintf("%v %v", key, value)))
	}

	return nil
}
