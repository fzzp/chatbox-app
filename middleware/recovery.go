package middleware

import (
	"chatbox-app/api"
	"chatbox-app/lib/errs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(*errs.ApiError)
				if ok {
					api.Fail(c, err)
				} else {
					// TODO: 其他错误
					err := fmt.Errorf(fmt.Sprintf("%v", err))
					api.Fail(c, errs.ErrServerError.AsException(err))
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
