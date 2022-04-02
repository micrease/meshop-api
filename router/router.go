package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/micrease/micrease-core/trace"
	"meshop-api/controller"
	"meshop-api/middleware"
)

func InitGinRouter() *gin.Engine {
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	ginRouter := gin.Default()
	//全局错误处理
	ginRouter.Use(middleware.RecoverError, middleware.RequestId(trace.TrafficKey))
	//每个controller一个分组,结构比较清晰
	prodGroup := ginRouter.Group("/demo/v1/product")
	{
		ctrl := controller.NewProduct()
		prodGroup.Handle("GET", "/list", ctrl.GetProductList)
		prodGroup.Handle("GET", "/detail", ctrl.GetProductDetail)
		prodGroup.Handle("POST", "/create", ctrl.CreateProduct)
		prodGroup.Handle("POST", "/update", ctrl.UpdateProduct)
		prodGroup.Handle("GET", "/delete", ctrl.DeleteProduct)
	}

	return ginRouter
}
