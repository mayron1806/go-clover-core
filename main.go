package main

import "github.com/mayron1806/go-fast/server"

func main() {
	f := server.NewFastServer(nil)
	f.Run()
}
