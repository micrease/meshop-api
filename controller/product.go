package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micrease/meshop-protos/product/pb"
	"github.com/micrease/micrease-core/errs"
	"github.com/micrease/micrease-core/structs"
	"github.com/micrease/micrease-core/validate"
	log "github.com/micro/go-micro/v2/logger"
	"meshop-api/common/rpcclient"
	"meshop-api/entity"
	"strconv"
)

type ProductController struct {
	BaseController
	productRpcService pb.ProductService
}

func NewProduct() *ProductController {
	ctrl := new(ProductController)
	return ctrl
}

/**
 * 获取查询列表
 */
func (this *ProductController) GetProductList(ctx *gin.Context) {

	//GET方式获取query string
	sizeStr := ctx.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeStr)
	//大写的原因,需要把小写的转义一下

	var prodReq pb.ProductRequest
	prodReq.Size = int32(size)

	prodResp, _ := rpcclient.Product.GetProductList(context.Background(), &prodReq)
	this.ResponseData(ctx, prodResp.Data)
}

/**
 * 获取详情
 */
func (this *ProductController) GetProductDetail(ctx *gin.Context) {

	sizeStr := ctx.DefaultQuery("id", "0")
	size, _ := strconv.Atoi(sizeStr)
	//大写的原因,需要把小写的转义一下
	var prodReq pb.Product
	prodReq.ProdId = int32(size)
	prodResp, _ := rpcclient.Product.GetProductDetail(context.Background(), &prodReq)
	log.Info("prodService.GetProductDetail:", prodResp)
	this.Response(ctx, prodResp.Data)
}

/**
 * 创建商品
 */
func (this *ProductController) CreateProduct(ctx *gin.Context) {
	//声明一个接收参数的实体
	var prodEntity entity.Product
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &prodEntity)
	//基它的判断，如果不成立，则抛出指定信息
	errs.PanicIfFalse(len(prodEntity.ProdName) < 4, errs.StatusCodeParamError, "商品名称不能小于4个字符")

	var prodReq pb.Product
	structs.CopyProperties(&prodEntity, &prodReq)

	prodResp, _ := rpcclient.Product.CreateProduct(context.Background(), &prodReq)
	this.Response(ctx, prodResp.Data)
}

/**
 * 更新商品
 */
func (this *ProductController) UpdateProduct(ctx *gin.Context) {
	//声明一个接收参数的实体
	var prodEntity entity.Product
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &prodEntity)
	//更新时必须要传
	errs.PanicIfFalse(prodEntity.ProdId <= 0, errs.StatusCodeParamError, "prodId参数不能为空")

	//grpc调用
	var prodReq pb.Product
	structs.CopyProperties(&prodEntity, &prodReq)

	prodResp, _ := rpcclient.Product.UpdateProduct(context.Background(), &prodReq)
	log.Info("prodService.UpdateProduct:", prodResp)

	//响应
	this.Response(ctx, prodResp.Data)
}

/**
 * 删除商品
 */
func (this *ProductController) DeleteProduct(ctx *gin.Context) {
	var prodReq pb.Product
	idStr := ctx.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(idStr)
	prodReq.ProdId = int32(id)

	//如果prodId小于等于0，则报错
	errs.PanicIfFalse(prodReq.ProdId <= 0, errs.StatusCodeParamError, "参数不正确")

	prodResp, _ := rpcclient.Product.DeleteProduct(context.Background(), &prodReq)
	log.Info("prodService.DeleteProduct:", prodResp)
	this.Response(ctx, prodResp.Data)
}
