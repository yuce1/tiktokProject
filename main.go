package main

import (
	"fmt"
	"os"

	"tiktok-go/repository"
	"tiktok-go/service"

	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()
	dsn_template := "host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf(
		dsn_template,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"))
	err := repository.InitDB(dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := gin.Default()
	r = initRouter(r)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
