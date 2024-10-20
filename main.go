package main

import "github.com/mayron1806/go-clover-core/server"

func main() {
	f := server.NewFastServer(nil)
	f.Run()
}
