package main

import (
	"log"
	"os"

	"github.com/sakul987/gObserver/pkg/api"
	"github.com/sakul987/gObserver/pkg/config"
)

func main(){
	os.Setenv("VITE_WS_URL", config.UI_VITE_WS_ADDR)
	os.Setenv("VITE_WS_PORT", config.UI_VITE_WS_PORT)
	
	if err := api.RunServer(); err != nil{
		log.Fatal(err)
	}
}