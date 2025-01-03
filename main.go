package main

import (
	"github.com/fulsep/docker-crud-backend/tree/main/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.CombineRouters(r)

	r.Run(":8888")
}
