package api

import (
	"chatbox-app/lib/errs"
	"log/slog"

	"github.com/gin-gonic/gin"
)

// JSON 响应数据
type JSON struct {
	BizCode   string `json:"bizCode"`             // 业务编码
	Message   string `json:"message"`             // 客户消息
	Data      any    `json:"data"`                // 任意数据
	Version   string `json:"version,omitempty"`   // 版本信息
	RequestID string `json:"requestId,omitempty"` // 请求ID，方便查询错误
}

func NewJSON(c *gin.Context, err *errs.ApiError, data any) *JSON {
	return &JSON{
		BizCode:   err.BizCode(),
		Message:   err.Message(),
		Data:      data,
		RequestID: c.Request.Header.Get("X-Request-ID"),
	}
}

func Succ(c *gin.Context, data any) {
	c.JSON(errs.ErrOK.StatusCode(), NewJSON(c, errs.ErrOK, data))
}

func Fail(c *gin.Context, err *errs.ApiError) {
	// log.Printf(
	// 	"Method: %s ->>> %s: %v\n",
	// 	c.Request.Method,
	// 	c.Request.URL.Path,
	// 	err.Error(),
	// )
	slog.ErrorContext(
		c,
		c.Request.Method+" ->> "+c.Request.URL.Path,
		slog.String("error", err.Error()),
	)
	c.JSON(errs.ErrOK.StatusCode(), NewJSON(c, err, nil))
}
