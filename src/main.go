package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/tubuarge/gora/server"
)

const Banner_File_Path = "asciiart.txt"

func banner() {
	b, err := os.ReadFile(Banner_File_Path)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(string(b))
}

func main() {
	banner()

	newRaftServer, err := server.NewRaft("./")
	if err != nil {
		log.Fatal(err)
	}
}
