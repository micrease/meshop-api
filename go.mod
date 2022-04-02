module meshop-api

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.1
	github.com/mattn/go-colorable v0.1.7
	github.com/micrease/meshop-protos v1.0.0
	github.com/micrease/micrease-core v1.0.5
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.27.0 // indirect
)

//解决etcdv3与grpc冲突问题https://www.jianshu.com/p/1971a27096b9
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
