package sstable

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"math/rand"
	"github.com/itsritiksingh/inMemoryStore/pkg/store"
)

var ssfolder = filepath.Join("pkg","sstable","tables")
var isMergeSSTablesRunning bool

func InitTable(s *store.Store) {
	tick := time.NewTicker(5 * time.Second)

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

		f.Write([]byte(fmt.Sprintf("%v %v\n", key, value)))
	}

	if !isMergeSSTablesRunning {
		isMergeSSTablesRunning = true
		go func() {
			mergeSSTables()
			isMergeSSTablesRunning = false
		}()
	}
	return nil
}

func mergeSSTables(){
	for {
		prefix:= "ss log*"
		matchingFiles , err := filepath.Glob(filepath.Join(ssfolder, prefix))
		if err != nil {
			log.Fatal("Error:", err)
			return
		}
		
		if len(matchingFiles) == 1 {
			return
		}
		
		for len(matchingFiles) > 1{
			_ , err := merge(matchingFiles[0],matchingFiles[1])
			if err != nil {
				fmt.Println(err.Error())
				fmt.Printf("not able to merge %s %s",matchingFiles[0],matchingFiles[1])
			}

			os.Remove(matchingFiles[0])
			os.Remove(matchingFiles[1])
			matchingFiles = matchingFiles[2:]
		}
	}
}

func merge(file1 string, file2 string) (string ,error){
	f1,err := os.OpenFile(file1,os.O_RDONLY,0644)
	if err != nil {
		return "",err
	}

	f2, err := os.OpenFile(file2,os.O_RDONLY,0644)
	if err != nil {
		return "", err
	}

	file3 := fmt.Sprintf("ss log merged%d%d%d",rand.Int(),time.Now().Day(),time.Now().Minute())
	f3, err := os.OpenFile(filepath.Join(ssfolder,file3),os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	
	if err != nil {
		return "", err
	}
	
	defer f1.Close()
	defer f2.Close()
	defer f3.Close()

	fileScanner1 := bufio.NewScanner(f1)
	fileScanner2 := bufio.NewScanner(f2)
	fileWriter3 := bufio.NewWriter(f3)

	fileScanner1.Split(bufio.ScanLines)
	fileScanner2.Split(bufio.ScanLines)
	
	s1 := fileScanner1.Scan()
	s2 := fileScanner2.Scan()

	for s1 || s2 {
		var line1 []string
		var line2 []string
		if s1 {
			line1 = strings.Split(fileScanner1.Text(), " ")
			s1 = fileScanner1.Scan()
		}

		if s2 {
			line2 = strings.Split(fileScanner2.Text()," ")
			s2 = fileScanner2.Scan()
		}


		if len(line2) == 0 || len(line1) > 0 && line1[0] < line2[0] {
			_, err := fileWriter3.WriteString(strings.Join(line1," ") +"\n")
			if err != nil {
				fmt.Printf("error while writing to merged log %s",err.Error())
			}

		} else {
			_, err := fileWriter3.WriteString(strings.Join(line2," ") +"\n")
			if err != nil {
				fmt.Printf("error while writing to merged log %s",err.Error())
			}
		}		
	}
	fileWriter3.Flush()

	return file3 , nil
}