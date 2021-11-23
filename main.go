package main

import (
	"deliveryhero/db"
	"deliveryhero/server"
)

func main() {
	db.Restore()
	db.BackupInterval(1)
	server.SetupServer()

}
