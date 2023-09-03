package valid

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans ut.Translator
)

func init() {
	InitTrans("zh")
}

// InitTrans 初始化翻译器
func InitTrans(Locale string) (err error) {
	//修改gin框架中的Validator引擎属性，实现自定制
	if V, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		V.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
			if name == "" {
				//没有label就用json
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			}
			if name == "-" {
				return ""
			}
			return name
		})

		//第一个参数是备用(fallback)的语言环境
		//后面的参数是应该支持的语言环境(支持多个)
		// uni := ut.New(zhT， zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)
		// Locale通常取决于http 请求头的' Accept-Language
		var ok bool
		//也可以使用uni. FindTranslator(...) 传入多个Locale进行查找
		trans, ok = uni.GetTranslator(Locale)
		if !ok {
			return fmt.Errorf("uni . GetTranslator(%s) failed", Locale)
		}

		//注册翻译器
		switch Locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(V, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(V, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(V, trans)
		}
		return
	}
	return
}

func Error(err error) (ret string) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	for _, e := range validationErrors {
		ret += e.Translate(trans) + ";"
	}
	return ret
}
