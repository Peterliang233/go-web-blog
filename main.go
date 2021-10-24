package main

import (
	"github.com/Peterliang233/go-blog/databases/mysql"
	"github.com/Peterliang233/go-blog/databases/redis"
	"github.com/Peterliang233/go-blog/router"
)

func main() {
	mysql.InitDB()
	redis.InitRedis()
	router.InitRouter()
}
