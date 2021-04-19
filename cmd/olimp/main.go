package main

import (
	"olimp/cmd/olimp/engine"
	"olimp/cmd/olimp/server"
)

func main() {
	srv := server.TServer{
		//s.Addr = www.example.com
		Addr:   ":8080",
		Engine: engine.CreateStore(),
	}

	srv.Run()
}
