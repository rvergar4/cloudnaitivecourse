// Launch microservice server- main.go
package main

import (
	"cloudnaitivecourse/lab12/microservice"
	"log"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
