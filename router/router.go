package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/gin/middleware"
	"github.com/micrease/micrease-core/trace"
	"meshop-api/handler"
)

func InitGinRouter() *gin.Engine {
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	ginRouter := gin.Default()
	//全局错误处理
	ginRouter.Use(middleware.Recover(true), middleware.RequestId(trace.TrafficKey))
	//每个controller一个分组,结构比较清晰
	prodGroup := ginRouter.Group("/demo/v1/product")
	{
		h := handler.NewProduct()
		prodGroup.Handle("GET", "/list", Handle(h.List))
		prodGroup.Handle("GET", "/detail", Handle(h.Detail))
		prodGroup.Handle("POST", "/create", Handle(h.Create))
		prodGroup.Handle("POST", "/update", Handle(h.Update))
		prodGroup.Handle("GET", "/delete", Handle(h.Delete))
	}
	return ginRouter
}

type HandleFunc func(ctx *context.Context)

func Handle(fn HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := new(context.Context)
		ctx.GinCtx = c
		fn(ctx)
	}
}
