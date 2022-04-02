package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/micrease/micrease-core/errs"
	"github.com/micrease/micrease-core/trace"
	sysConfig "meshop-api/config"
	"net/http"
	"strconv"
	"strings"
)

func RecoverError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {

			if c.IsAborted() {
				c.Status(200)
			}

			me := errs.NewError()
			me.Code = 500
			switch errStr := err.(type) {
			case string:
				//格式如:5001#message
				p := strings.Split(errStr, errs.ERR_DS)
				l := len(p)
				if l >= 2 {
					//501#message
					statusCode, e := strconv.Atoi(p[0])
					if e != nil {
						break
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
					me.TraceId = c.Request.Header.Get(trace.TrafficKey)
					c.JSON(http.StatusOK, me)
				} else {
					//格式如:message
					c.JSON(http.StatusOK, gin.H{
						"code": 5000,
						"msg":  errStr,
					})
				}
			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
