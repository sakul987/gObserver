package main

import (
	"log"
	"github.com/sakul987/gObserver/pkg/api"
)

func main(){
	if err := api.RunServer(); err != nil{
		log.Fatal(err)
	}
}