package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mini-tiktok/dao"
)

func main() {
	r := gin.Default()

	initRouter(r)
	dao.Init()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
