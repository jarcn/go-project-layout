package main

import (
	"github.com/gin-gonic/gin"
	"kupu.com/chenjia/p/api"
)

func main() {
	r := gin.Default()
	api.Register(r)
	r.Run(":8080")
}
