package main

import (
	"deliveryhero/db"
	"deliveryhero/server"
	"time"
)

func main() {
	go func() {
		db.Restore()
		db.BackupInterval(10 * time.Minute)
	}()
	server.SetupServer()
}
