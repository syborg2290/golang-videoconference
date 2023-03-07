package main

import (
	"log"
	"video-conference/internal/server"
)

func main(){
	if err := server.Run(); err != nil{
		log.Fatalln(err.Error())
	}
}