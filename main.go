package main

import (
	"deliveryhero/server"
)

func main() {
	// go db.Restore()
	// go db.BackupInterval(1 * time.Minute)
	server.SetupServer()
}
