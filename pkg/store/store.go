package store

import (
	"errors"
	"fmt"
	"time"
	"sync"

	"github.com/itsritiksingh/inMemoryStore/pkg/wal"
)

type Store struct {
	Store map[string]string
	wal *wal.Wal
	Mu  *sync.RWMutex
}

func Init() *Store {
	w := wal.Init()
	return &Store{
		Store: make(map[string]string),
		wal: w,
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

func (s *Store) Put(key string,value string) (bool,error){
	s.Mu.Lock()
	defer s.Mu.Unlock()
	_ , isFound := s.Store[key]

	if isFound {
		return false, errors.New("key already exist")
	}
	s.Store[key] = value
	s.wal.Write([]byte(fmt.Sprintf("%v %v %v %v",key,value,false,time.Now().UnixMicro())))
	return true,nil
}

func (s *Store) Upsert(key string,value string) (bool , error){
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Store[key] = value
	s.wal.Write([]byte(fmt.Sprintf("%v %v %v %v",key,value,false,time.Now().UnixMicro())))
	return true , nil
}

func (s *Store) Delete(key string,value string) (bool,error){
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Store,key)
	s.wal.Write([]byte(fmt.Sprintf("%v %v %v %v",key,value,true,time.Now().UnixMicro())))
	return true, nil
}