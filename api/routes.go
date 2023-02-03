package api

import (
	"github.com/gin-gonic/gin"
	"kupu.com/chenjia/p/pkg/source/http_spec"
)

func Register(r *gin.Engine) {
	// 初始页面
	r.GET("/welcome", http_spec.Welcome)
}
