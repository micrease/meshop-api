package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/trace"
	"meshop-api/handler"
	"meshop-api/middleware"
)

func InitGinRouter() *gin.Engine {
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	ginRouter := gin.Default()
	//全局错误处理
	ginRouter.Use(middleware.Recover, middleware.RequestId(trace.TrafficKey))
	//每个controller一个分组,结构比较清晰
	prodGroup := ginRouter.Group("/demo/v1/product")
	{
		ctrl := handler.NewProduct()
		prodGroup.Handle("GET", "/list", Handle(ctrl.List))
		prodGroup.Handle("GET", "/detail", Handle(ctrl.Detail))
		prodGroup.Handle("POST", "/create", Handle(ctrl.Create))
		prodGroup.Handle("POST", "/update", Handle(ctrl.Update))
		prodGroup.Handle("GET", "/delete", Handle(ctrl.Delete))
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
