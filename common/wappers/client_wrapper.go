package wappers

import (
	"context"
	"github.com/asim/go-micro/v3/client"
	micro_errors "github.com/asim/go-micro/v3/errors"
	log "github.com/asim/go-micro/v3/logger"
)

//定义一个wrapper,继承Client,重写Call方法
type ClientWrapper struct {
	client.Client
}

func NewClientWrapper(c client.Client) client.Client {
	return &ClientWrapper{c}
}

/**
 * 在rpc调用时统一做一些通用操作，比如log,超时熔断等
 */
func (w *ClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Info("ClientWraper.Call,req:", req.Body())
	err := w.Client.Call(ctx, req, rsp)
	if err != nil {
		panic(micro_errors.Parse(err.Error()))
	}
	return err
}
