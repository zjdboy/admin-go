package main

import (
	"admin-go/apis"
)

func main() {
	router := apis.Apis()
	router.Run(":8888")
}
