package main

import "fmt"

type server struct {
	h map[string]func(key, value string)
}

func (s *server) Register(name string, h func(key, value string)) {
	if s.h[name] != nil {
		delete(s.h, name)
	}
	s.h[name] = h
}

func (s *server) Run(name, key, value string) {
	if s.h[name] != nil {
		s.h[name](key, value)
	}
}

var Server = &server{
	h: make(map[string]func(key, value string)),
}

func client(key, value string) {
	fmt.Printf("%s %s\n", key, value)
}

func client2(key, value string) {
	fmt.Printf("%s %s\n", key, value)
}

func main(){
	Server.Register("client", client)
	Server.Register("client2", client2)

	Server.Register("client3", func(key, value string) {
		fmt.Printf("%s %s\n", key, value)
	})
	Server.Run("client", "key", "value")
	Server.Run("client2", "key1", "value2")
	Server.Run("client3", "key1", "value2")
}
