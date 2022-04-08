package handler

import (
	"context"
	"github.com/micrease/meshop-protos/product/pb"
	mctx "github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/errs"
	"github.com/micrease/micrease-core/gin"
	"github.com/micrease/micrease-core/structs"
	"github.com/micrease/micrease-core/validate"
	log "github.com/micro/go-micro/v2/logger"
	"meshop-api/common/rpcclient"
	"meshop-api/entity"
	"strconv"
)

type Product struct {
	gin.Handler
	productRpcService pb.ProductService
}

func NewProduct() *Product {
	ctrl := new(Product)
	return ctrl
}

/**
 * 获取查询列表
 */
func (this *Product) List(ctx *mctx.Context) {
	//GET方式获取query string
	sizeStr := ctx.GinCtx.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeStr)
	//大写的原因,需要把小写的转义一下
	var prodReq pb.ProductRequest
	prodReq.Size = int32(size)
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := rpcclient.Product.GetProductList(context.Background(), &prodReq)
	this.ResponseData(ctx, prodResp.Data)
}

/**
 * 获取详情
 */
func (this *Product) Detail(ctx *mctx.Context) {
	sizeStr := ctx.GinCtx.DefaultQuery("id", "0")
	size, _ := strconv.Atoi(sizeStr)
	//大写的原因,需要把小写的转义一下
	var prodReq pb.Product
	prodReq.ProdId = int32(size)
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := rpcclient.Product.GetProductDetail(context.Background(), &prodReq)
	log.Info("prodService.GetProductDetail:", prodResp)
	this.Response(ctx, prodResp.Data)
}

/**
 * 创建商品
 */
func (this *Product) Create(ctx *mctx.Context) {
	//声明一个接收参数的实体
	var prodEntity entity.Product
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &prodEntity)
	//基它的判断，如果不成立，则抛出指定信息
	errs.PanicIf(len(prodEntity.ProdName) < 4, errs.StatusParamError, "商品名称不能小于4个字符")
	var prodReq pb.Product
	structs.Copy(&prodEntity, &prodReq)
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := rpcclient.Product.CreateProduct(context.Background(), &prodReq)
	this.Response(ctx, prodResp.Data)
}

/**
 * 更新商品
 */
func (this *Product) Update(ctx *mctx.Context) {
	//声明一个接收参数的实体
	var prodEntity entity.Product
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &prodEntity)
	//更新时必须要传
	errs.PanicIf(prodEntity.ProdId <= 0, errs.StatusParamError, "prodId参数不能为空")
	//grpc调用
	var prodReq pb.Product
	structs.Copy(&prodEntity, &prodReq)
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := rpcclient.Product.UpdateProduct(context.Background(), &prodReq)
	log.Info("prodService.UpdateProduct:", prodResp)
	//响应
	this.Response(ctx, prodResp.Data)
}

/**
 * 删除商品
 */
func (this *Product) Delete(ctx *mctx.Context) {
	var prodReq pb.Product
	idStr := ctx.GinCtx.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(idStr)
	prodReq.ProdId = int32(id)
	//如果prodId小于等于0，则报错
	errs.PanicIf(prodReq.ProdId <= 0, errs.StatusParamError, "参数不正确")
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := rpcclient.Product.DeleteProduct(context.Background(), &prodReq)
	log.Info("prodService.DeleteProduct:", prodResp)
	this.Response(ctx, prodResp.Data)
}
