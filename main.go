package main

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"./raft"
)

const Banner_File_Path = "asciiart.txt"

func banner() {
	b, err := ioutil.ReadFile(Banner_File_Path)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(string(b))
}

func main() {
	banner()

	newRaftServer, err := raft.NewRaft("./")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newRaftServer)
}
