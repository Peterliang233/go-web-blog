package main

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/router"
)

func main() {
	databases.InitDB()
	router.InitRouter()
}
