package store

import (
	"errors"
)

type Store struct {
	store map[string]string
}

func Init() *Store {
	return &Store{
		store: make(map[string]string),
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
	return true,nil
}

func (s *Store) Upsert(key string,value string) (bool , error){
	s.store[key] = value
	return true , nil
}

func (s *Store) Delete(key string,value string) (bool,error){
	delete(s.store,key)
	return true, nil
}