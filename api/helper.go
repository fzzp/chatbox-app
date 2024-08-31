package api

import (
	"chatbox-app/config"
	"chatbox-app/lib/errs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"strings"

	"github.com/go-playground/validator/v10"
)

func GenHashPassword(plainText string) (string, error) {
	pswdHash, err := bcrypt.GenerateFromPassword([]byte(plainText), 12)
	if err != nil {
		return "", err
	}
	return string(pswdHash), nil
}

func CheckPassword(plainText, pswdHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(pswdHash), []byte(plainText))
}

// func ShouldBindJSON(c *gin.Context, input any) bool {
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		Fail(c, errs.ErrBadRequest.AsException(err))
// 		return false
// 	}
// 	return true
// }

// 绑定参数类型
type BindType string

var (
	BindJson  BindType = "JSON"
	BindUri   BindType = "URI"
	BindQuery BindType = "QUERY"
)

// ShouldBind 绑定请求参数，对 gin ShouldBindXXX 进行二次封装，针对 struct 的 form、json、uri tag
// 如果 bindTypes 不传则默认是 BindJson
func ShouldBind(c *gin.Context, req interface{}, bindTypes ...BindType) bool {
	var err error

	// 设置默认值
	bindType := BindJson
	if len(bindTypes) > 0 {
		bindType = bindTypes[0]
	}
	switch bindType {
	case BindJson:
		err = c.ShouldBindJSON(req)
	case BindQuery:
		err = c.ShouldBindQuery(req)
	case BindUri:
		err = c.ShouldBindUri(req)
	default:
		err = c.ShouldBind(req)
	}

	if err != nil {
		return handleError(c, err)
	}

	return true
}

// handleError 处理 Bind 请求参数发生的错误
func handleError(c *gin.Context, err error) bool {
	// 断言错误是否为 validator/v10 的验证错误信息
	verrs, ok := err.(validator.ValidationErrors)
	if !ok { // 其他方面的参数不匹配
		Fail(c, errs.ErrBadRequest.AsException(err))
		return false
	}

	// 对错误信息进行翻译, 得到的是 map[string]string 结构数据
	merrs := verrs.Translate(config.App.Trans)

	var msgs []string
	for _, e := range merrs {
		// 判断是否是自定义错误
		if strings.Contains(e, "vPhone") {
			msgs = append(msgs, "手机号码不正确")
		} else {
			msgs = append(msgs, e)
		}
	}

	Fail(c, errs.ErrBadRequest.AsMessage(strings.Join(msgs, ";")))

	return false
}
