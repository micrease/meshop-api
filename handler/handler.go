package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micrease/micrease-core/errs"
)

type GinHandler struct {
	errs.Error
}

func (ctrl GinHandler) ResponseData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{"status": 200, "message": "操作成功", "data": data})
}

func (ctrl GinHandler) Success(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": 200, "message": "操作成功", "data": ""})
}

func (ctrl GinHandler) Response(ctx *gin.Context, data interface{}) {
	ctrl.ResponseData(ctx, data)
}
