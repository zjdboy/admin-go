package main

import (
	"admin-go/handlers"
)

func main() {
	router := handlers.Handlers()
	router.Run(":8888")
}
