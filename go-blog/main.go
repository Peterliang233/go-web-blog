package main

import (
	"github.com/Peterliang233/go-blog/model"
	"github.com/Peterliang233/go-blog/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}