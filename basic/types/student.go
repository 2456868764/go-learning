package main

import (
	"errors"
	"fmt"
)

var _ Study = (*study)(nil)

type Study interface {
	Listen(msg string) string
	Speak(msg string) string
	Read(msg string) string
	Write(msg string) string
}

type study struct {
	Name string
}

func (s *study) Listen(msg string) string {
	return s.Name + " 听 " + msg
}

func (s *study) Speak(msg string) string {
	return s.Name + " 说 " + msg
}

func (s *study) Read(msg string) string {
	return s.Name + " 读 " + msg
}

func (s *study) Write(msg string) string {
	return s.Name + " 写 " + msg
}

func New(name string) (Study, error) {
	if name == "" {
		return nil, errors.New("name required")
	}

	return &study{
		Name: name,
	}, nil
}

func main() {
	name := "Xiao Ming "
	s, err := New(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}
