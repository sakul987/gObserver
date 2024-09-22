package main

import (
	"fmt"
	"log"
	"github.com/sakul987/gObserver/pkg/api"
)

func main(){
	fmt.Println("init")
	
	if err := api.RunServer(); err != nil{
		log.Fatal(err)
	}
}