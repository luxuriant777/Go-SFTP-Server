package main

import (
	"log"
	"sftp-server/cmd/server"
)

func main() {
	err := server.Run("/path/to/your/config.json")
	if err != nil {
		log.Fatal(err)
	}
}
