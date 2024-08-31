package errs

import "net/http"

var (
	ErrOK                  = NewApiError(http.StatusOK, "200", "请求成功")
	ErrBadRequest          = NewApiError(http.StatusBadRequest, "400", "入参错误")
	ErrUnauthorized        = NewApiError(http.StatusUnauthorized, "401", "请先登陆")
	ErrForbidden           = NewApiError(http.StatusForbidden, "403", "禁止访问")
	ErrNotFound            = NewApiError(http.StatusNotFound, "404", "查无此记录")
	ErrMethodNotAllowed    = NewApiError(http.StatusMethodNotAllowed, "405", "请求方法不支持")
	ErrRecordExists        = NewApiError(http.StatusConflict, "409", "数据重复")
	ErrUnprocessableEntity = NewApiError(http.StatusUnprocessableEntity, "422", "请求无法处理")
	ErrTooManyRequests     = NewApiError(http.StatusTooManyRequests, "429", "请求繁忙")
	ErrServerError         = NewApiError(http.StatusInternalServerError, "500", "请求错误，请稍后重试！")
)
