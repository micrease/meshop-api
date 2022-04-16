package remote_service

import (
	"github.com/asim/go-micro/v3"
	"github.com/micrease/meshop-protos/product/pb"
	"meshop-api/common/wappers"
	sysConfig "meshop-api/config"
)

const (
	//服务名:对应注册中心服务名,a要访问b服务的接口，这里是b服务的服务名
	remoteServiceProduct = "meshop-product-service"
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
func Register() {
	clientServiceName := sysConfig.Get().Service.ServiceName
	//创建一个rpc调用client
	apiClient := micro.NewService(
		micro.Name(clientServiceName),
		//添加一个wrapper,对Call方法重写,前置一些统一的操作逻辑
		micro.WrapClient(wappers.NewClientWrapper),
	)
	//"product-service" 对应服务注册名和注册中心product-service服务名称保持一致micro.rpc-service.product
	//TODO::添加需要使用的远程服务
	Product = pb.NewProductService(remoteServiceProduct, apiClient.Client())
}
