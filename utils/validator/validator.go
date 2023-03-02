package validator

import (
	"fmt"
	"ginVueBlog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (string, int) { //泛型
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New()) //翻译实例化，转换为zh_Hans_CN
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans) //对翻译注册
	if err != nil {
		fmt.Println("err:", err)
	}

	//label值取代字段
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data) //验证结构公开的字段，并自动验证嵌套结构
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) { //错误执行
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
