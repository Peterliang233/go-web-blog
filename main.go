package main

import (
	"github.com/Peterliang233/go-blog/router"
	"github.com/Peterliang233/go-blog/service/v1/model"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
