package main

import "github.com/sandeepraju/blockchain/api"

func main() {
	server := api.NewServer()
	server.Start()
}
