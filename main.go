package main

import "github.com/mproyyan/gin-rest-api/cmd"

func main() {
	server := cmd.Boot()
	server.Run()
}
