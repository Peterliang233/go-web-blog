package main

import (
	"github.com/Peterliang233/go-blog/model"
	"github.com/Peterliang233/go-blog/router"
)

func main() {
	model.InitDB()
	router.InitRouter()
}
