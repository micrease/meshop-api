package controller

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (ctrl BaseController) ResponseFailed(ctx *gin.Context, code int32, message string) {
	ctx.JSON(200, gin.H{"status": code, "message": message})
}

func (ctrl BaseController) ResponseData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{"status": 200, "message": "操作成功", "data": data})
}

func (ctrl BaseController) ResponseMessage(ctx *gin.Context, code int32, message string) {
	ctx.JSON(200, gin.H{"status": code, "message": message, "data": ""})
}

func (ctrl BaseController) ResponseSuccess(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": 200, "message": "操作成功", "data": ""})
}

func (ctrl BaseController) ResponseRPCNilFailed(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": 5000, "message": "call rpc service error,nil result"})
}

func (ctrl BaseController) Response(ctx *gin.Context, data interface{}) {
	ctrl.ResponseData(ctx, data)
}
