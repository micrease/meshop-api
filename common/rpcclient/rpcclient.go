package rpcclient

import (
	"github.com/micrease/meshop-protos/product/pb"
	"github.com/micro/go-micro/v2"
	"meshop-api/common/wappers"
	sysConfig "meshop-api/config"
)

/**
 * 统一管理RPC Client,在服务启动时创建一次
 */
var (
	Product pb.ProductService
)

/**
 * 把需要远程调用的client都创建在这里
 */
func RegisterRpcClient() {
	clientServiceName := sysConfig.Get().Service.ServiceName
	//创建一个rpc调用client
	apiClient := micro.NewService(
		micro.Name(clientServiceName),
		//添加一个wrapper,对Call方法重写,前置一些统一的操作逻辑
		micro.WrapClient(wappers.NewClientWrapper),
	)
	//"product-service" 对应服务注册名和注册中心product-service服务名称保持一致micro.rpc-service.product
	Product = pb.NewProductService("meshop-product-service", apiClient.Client())
}
