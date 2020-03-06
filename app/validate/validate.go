/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:57:13
 * @LastEditTime: 2020-03-05 16:02:30
 * @FilePath: /XGBlog/app/validate/validate.go
 * @Description:
 */

package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/zh"
)

// Validate 认证结构体
type Validate struct {
	validator *validator.Validate
	trans     ut.Translator
	errs      validator.ValidationErrors
}

// Default 默认方法
func Default() (*Validate, error) {
	validate := &Validate{}
	validate.validator = validator.New()
	lang := zh.New()
	uni := ut.New(lang, lang)
	trans, _ := uni.GetTranslator("zh")
	validate.SetTrans(trans)
	err := validate.registerDefaultTranslations(validate.trans)
	return validate, err
}

// SetTrans 设置翻译
func (v *Validate) SetTrans(trans ut.Translator) {
	v.trans = trans
}

// registerDefaultTranslations 注册默认翻译器
func (v *Validate) registerDefaultTranslations(trans ut.Translator) error {
	return en_translations.RegisterDefaultTranslations(v.validator, trans)
}

// CheckStruct 检查结构体
func (v *Validate) CheckStruct(s interface{}) bool {
	err := v.validator.Struct(s)
	if err != nil {
		v.errs = err.(validator.ValidationErrors)
		return false
	}
	return true
}

// GetAllError 获取所有错误
func (v *Validate) GetAllError() []string {
	var errList []string
	for _, e := range v.errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(v.trans))
	}
	return errList
}

// GetOneError 获取一个错误
func (v *Validate) GetOneError() string {
	for _, e := range v.errs {
		// can translate each error one at a time.
		return e.Translate(v.trans)
	}
	return ""
}
