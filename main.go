package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	initRouter(r)

	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
