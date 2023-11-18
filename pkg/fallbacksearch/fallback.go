package fallbacksearch

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/itsritiksingh/inMemoryStore/pkg/wal"
)

var (
	KeyNotFound = errors.New("keynotfound")
	KeyDeleted = errors.New("keydeleted")
)

func search(key string) (string, error) {
	res , err := searchWal(key)
	if errors.Is(err,KeyDeleted){
		return "" ,KeyDeleted
	}

	if res != "" {
		return res , nil
	}

	return "",nil
}

func searchWal(key string) (string , error) {
	lock := &sync.RWMutex{}
	lock.Lock()
	defer lock.Unlock()
	
	file, err := os.Open(wal.WalFileName)
	if err != nil {
		return "",err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "",err
	}

	fileSize := fileInfo.Size()
	buf := make([]byte, 1)

	for offset := int64(1); offset <= fileSize; offset++ {
		file.Seek(-offset, 2) // Seek from the end of the file
		file.Read(buf)

		if buf[0] == '\n' {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				s := scanner.Text()
				arr := strings.Split(s," ")
				if arr[0] == key {
					if isDeleted,_ := strconv.ParseBool(arr[2]); isDeleted {
						return "",KeyDeleted
					}

					return arr[1] , nil
				}
			}
		}

	}
	return "", KeyNotFound
}
