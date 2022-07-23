package main

import (
	"log"
	"encoding/json"
	lc "github.com/ibrahim925/LogiCore"
)

func main() {
	client, err := lc.NewClient("https://garmindev.dev.logisensebilling.com", "admin", "admin", "044b8ad6006845c29446b2f18e5b5909")
	if err != nil {
		log.Panicln(err)
		return 
	}
	
	service, err := client.GetService()
	if err != nil {
		log.Panicln(err)
		return
	}

	body, err := json.MarshalIndent(service, "", "    ") 
	log.Println(body)
}
