package store

import (
	"errors"
	"fmt"
	"time"

	"github.com/itsritiksingh/inMemoryStore/pkg/wal"
)

type Store struct {
	store map[string]string
	wal *wal.Wal
}

func Init() *Store {
	w := wal.Init()
	return &Store{
		store: make(map[string]string),
		wal: w,
	}
}

func (s *Store) Get(key string) (string, error) {
	val, isFound := s.store[key]

	if !isFound {
		return "", errors.New("key not found")
	}
	return val, nil
}

func (s *Store) Put(key string,value string) (bool,error){
	_ , isFound := s.store[key]

	if isFound {
		return false, errors.New("key already exist")
	}
	s.store[key] = value
	s.wal.Write([]byte(fmt.Sprintf("%v %v %v %v",key,value,false,time.Now().UnixMicro())))
	return true,nil
}

func (s *Store) Upsert(key string,value string) (bool , error){
	s.store[key] = value
	s.wal.Write([]byte(fmt.Sprintf("%v %v %v %v",key,value,false,time.Now().UnixMicro())))
	return true , nil
}

func (s *Store) Delete(key string,value string) (bool,error){
	delete(s.store,key)
	s.wal.Write([]byte(fmt.Sprintf("%v %v %v %v",key,value,true,time.Now().UnixMicro())))
	return true, nil
}