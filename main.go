package main

import (
	"mini-tiktok/dao"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitDeps() {
	dao.Init()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

}
func main() {
	r := gin.Default()
	initRouter(r)
	InitDeps()
	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
