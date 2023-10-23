package store

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/itsritiksingh/inMemoryStore/pkg/wal"
)

type Store struct {
	Store map[string]string
	Mu  *sync.RWMutex
}

func Init() *Store {
	return &Store{
		Store: make(map[string]string),
		Mu: &sync.RWMutex{},
	}
}

func (s *Store) Get(key string) (string, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	val, isFound := s.Store[key]

	if !isFound {
		return "", errors.New("key not found")
	}
	return val, nil
}

func (s *Store) GetAllKeys() ([]string){
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	keys := make([]string, 0, len(s.Store))
    for k := range s.Store {
        keys = append(keys, k)
    }

	return keys
}

func (s *Store) Put(key string,value string) (bool,error){
	s.Mu.Lock()
	defer s.Mu.Unlock()
	_ , isFound := s.Store[key]

	if isFound {
		return false, errors.New("key already exist")
	}
	s.Store[key] = value
	wal.Write([]byte(fmt.Sprintf("%v %v %v %v\n",key,value,false,time.Now().UnixMicro())))
	return true,nil
}

func (s *Store) Upsert(key string,value string) (bool , error){
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Store[key] = value
	wal.Write([]byte(fmt.Sprintf("%v %v %v %v\n",key,value,false,time.Now().UnixMicro())))
	return true , nil
}

func (s *Store) Delete(key string,value string) (bool,error){
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Store,key)
	wal.Write([]byte(fmt.Sprintf("%v %v %v %v\n",key,value,true,time.Now().UnixMicro())))
	return true, nil
}