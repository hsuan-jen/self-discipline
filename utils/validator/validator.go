package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func Verify(c *gin.Context, params interface{}) (bool, string) {

	zhTrans := zh.New()
	enTrans := en.New()

	uni := ut.New(zhTrans, enTrans) // 创建一个通用转换器
	curLocales := "zh"              // 设置当前语言类型
	var trans ut.Translator         // 获取对应语言的转换器

	validate := validator.New()

	switch curLocales {
	case "en":
		trans, _ = uni.GetTranslator(curLocales) // 获取对应语言的转换器
		_ = en_trans.RegisterDefaultTranslations(validate, trans)

	case "zh":
		fallthrough
	default:
		trans, _ = uni.GetTranslator(curLocales) // 获取对应语言的转换器
		// 内置tag注册 中文翻译器
		_ = zh_trans.RegisterDefaultTranslations(validate, trans)

		// 注册 RegisterTagNameFunc
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	err := validate.Struct(params)
	if err != nil {

		err := err.(validator.ValidationErrors)
		for _, e := range err {

			return false, e.Translate(trans)
		}
	}

	return true, ""
}
