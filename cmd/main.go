package main

import (
	"github.com/sakul987/gObserver/pkg/api"
	"log"
)

func main() {
	if err := api.RunServer(); err != nil {
		log.Fatal(err)
	}
}
