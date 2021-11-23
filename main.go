package main

import (
	"deliveryhero/db"
	"deliveryhero/server"
	"time"
)

func main() {
	go func() {
		db.Restore()
		db.BackupInterval(1 * time.Minute)
	}()
	server.SetupServer()
}
