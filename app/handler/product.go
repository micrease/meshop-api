package handler

import (
	"context"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/micrease/meshop-protos/product/pb"
	me "github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/errs"
	"github.com/micrease/micrease-core/gin"
	"github.com/micrease/micrease-core/validate"
	"meshop-api/app/service/remote_service"
)

type Product struct {
	gin.Handler
}

func NewProduct() *Product {
	return &Product{}
}

/**
 * 获取查询列表
 */
func (this *Product) List(ctx *me.Context) {
	//大写的原因,需要把小写的转义一下
	var req pb.ProductPageReq
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &req)
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := remote_service.Product.PageList(context.Background(), &req)
	this.ResponseData(ctx, prodResp.Data)
}

/**
 * 获取详情
 */
func (this *Product) Detail(ctx *me.Context) {
	var req pb.ProductDetailReq
	//remote_service远程调用,如果是单体项目可以用service代替
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := remote_service.Product.Detail(context.Background(), &req)
	log.Info("prodService.GetProductDetail:", prodResp)
	this.Response(ctx, prodResp.Data)
}

/**
 * 创建商品
 */
func (this *Product) Create(ctx *me.Context) {
	//声明一个接收参数的实体
	var req pb.ProductInsertReq
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &req)
	//基它的判断，如果不成立，则抛出指定信息
	errs.PanicIf(len(req.Name) < 4, errs.StatusParamError, "商品名称不能小于4个字符")
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := remote_service.Product.Create(context.Background(), &req)
	this.Response(ctx, prodResp.Data)
}

/**
 * 更新商品
 */
func (this *Product) Update(ctx *me.Context) {
	//声明一个接收参数的实体
	var req pb.ProductUpdateReq
	//绑定参数，并验证合法性
	validate.BindWithPanic(ctx, &req)
	//更新时必须要传
	errs.PanicIf(req.Id <= 0, errs.StatusParamError, "prodId参数不能为空")
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := remote_service.Product.Update(context.Background(), &req)
	log.Info("prodService.UpdateProduct:", prodResp)
	//响应
	this.Response(ctx, prodResp.Data)
}

/**
 * 删除商品
 */
func (this *Product) Delete(ctx *me.Context) {
	var req pb.ProductDeleteReq
	//如果prodId小于等于0，则报错
	errs.PanicIf(req.Id <= 0, errs.StatusParamError, "参数不正确")
	//此处无须处理err,err已经在client_warpper中拦截了
	prodResp, _ := remote_service.Product.Delete(context.Background(), &req)
	log.Info("prodService.DeleteProduct:", prodResp)
	this.Response(ctx, prodResp.Data)
}
