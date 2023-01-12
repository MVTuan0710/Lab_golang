package main

import (
	"myapp/routers"
	"myapp/storage"
)

func main() {
	// Echo instance
	storage.NewDB()
	// Routes
	routers.Api()
}
