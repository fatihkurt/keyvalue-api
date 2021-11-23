package main

import (
	"deliveryhero/db"
	"deliveryhero/server"
	"time"
)

func main() {
	go db.Restore()
	go db.BackupInterval(1 * time.Minute)
	server.SetupServer()
}
