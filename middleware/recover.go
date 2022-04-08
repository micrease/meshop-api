package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micrease/micrease-core/errs"
	"github.com/micrease/micrease-core/trace"
	micro_errors "github.com/micro/go-micro/v2/errors"
	sysConfig "meshop-api/config"
	"net/http"
	"strconv"
	"strings"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if c.IsAborted() {
				c.Status(200)
			}
			switch errStr := err.(type) {
			case *micro_errors.Error:
				str := fmt.Sprintf("%s%s%s", errStr.Id, errs.ERR_DS, errStr.Detail)
				me := GetError(str)
				me.TraceId = c.Request.Header.Get(trace.TrafficKey)
				c.JSON(http.StatusOK, me)
			case string:
				//格式如:5001#message
				me := GetError(errStr)
				me.TraceId = c.Request.Header.Get(trace.TrafficKey)
				c.JSON(http.StatusOK, me)
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}

func GetError(str string) errs.Error {
	me := errs.Error{}
	me.Code = 500
	p := strings.Split(str, errs.ERR_DS)
	l := len(p)
	if l >= 2 {
		//501#message
		statusCode, e := strconv.Atoi(p[0])
		if e != nil {
			me.Error = e.Error()
			me.Message = e.Error()
			return me
		}
		me.Code = statusCode
		me.Message = p[1]
		conf := sysConfig.Get()
		//debug模式显示错误信息
		if conf.Service.DebugEnable {
			//Error
			if l >= 3 {
				me.Error = p[2]
			}

			//traceInfo
			if l >= 4 {
				me.Error = me.Error + "," + p[3]
			}
		}
	} else {
		me.Message = str
	}
	return me
}
