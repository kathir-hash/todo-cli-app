package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{filename: filename}
}


func (s *Storage[T]) Save(data *T) error {
	filedata, err := json.MarshalIndent(data,"","    ")
	if err != nil {
		return  err
	}
	return os.WriteFile(s.filename,filedata,0644)
}

func (s *Storage[T]) Load(data *T) error {
	filedata ,err := os.ReadFile(s.filename)
	if err != nil {
		return  err
	}
	return json.Unmarshal(filedata,data)
}