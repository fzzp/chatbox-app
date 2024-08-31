package validate

import (
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	ut "github.com/go-playground/universal-translator"
)

var vPhoneRX = regexp.MustCompile(`^1[3456789]\d{9}$`)

// NewValidation 设置翻译器，locale 为:zh 或 en
func NewValidation(locale string) (trans ut.Translator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册手机号码自定义验证函数和tag
		v.RegisterValidation("vPhone", func(fieldLevel validator.FieldLevel) bool {
			if val, ok := fieldLevel.Field().Interface().(string); ok {
				if vPhoneRX.MatchString(val) {
					return true
				}
			}
			return false
		})

		// 注册一个函数获取字段用于翻译的tag，这里仅获取与binding库相关的三个tag。
		// 也可以自定义描述tag 如: field string `label:"字段"`
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 1. 优先处理 json tag
			// filed string `json:"aa,bb,cc"` -> aa
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			if len([]rune(name)) > 0 {
				return name
			}
			// 2. 如果 json tag 没有则处理 form tag
			name = strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			if len([]rune(name)) > 0 {
				return name
			}
			// 3. 最后处理 uri tag
			name = strings.SplitN(field.Tag.Get("uri"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 获取翻译器
		zhTrans := zh.New()
		enTrans := en.New()
		// 第一个参数时回退翻译（备用），后面是支持翻译器
		uni := ut.New(enTrans, zhTrans, enTrans)
		// 获取指定的翻译器，如果不存在则使用备用
		var found bool
		trans, found = uni.GetTranslator(locale)
		if !found {
			log.Panicf(`uni.GetTranslator(%q) not found`, locale)
			return trans
		}

		// 设置默认翻译器
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
