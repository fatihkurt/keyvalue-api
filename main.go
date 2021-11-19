package main

import (
	"deliveryhero/server"

	"github.com/fvbock/endless"
)

func main() {
	router := server.SetupRouter()
	err := endless.ListenAndServe(":5000", router)

	if err != nil {
		panic(err)
	}
}
