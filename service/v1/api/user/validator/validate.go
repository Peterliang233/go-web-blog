package validator

import (
	"fmt"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	uniTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	TransZh "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// 进行数据验证并且指定特定的语言进行错误的响应
func Validate(data interface{}) (string, int) {
	validate := validator.New()                 // 实例化
	uni := uniTrans.New(zh_Hans_CN.New())       // 实例化
	trans, _ := uni.GetTranslator("zh_Hans_CN") // 指定翻译成的语言
	// 验证器注册翻译器
	err := TransZh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data) //进行数据验证，其实就是一个反射的过程
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.Error
		}
	}
	return "验证成功", errmsg.Success
}
