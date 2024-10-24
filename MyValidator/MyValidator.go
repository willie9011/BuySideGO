package MyValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//要先安裝 go get github.com/go-playground/validator/v10

func MyUrlValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	// 验证字段值是否符合要求
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,}$`)
	return re.MatchString(value)
}
