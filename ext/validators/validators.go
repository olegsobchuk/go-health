package validators

import (
	"reflect"

	"github.com/olegsobchuk/go-health/configs"
	validator "gopkg.in/go-playground/validator.v8"
)

// UniqEmail add unique validation to field *email*
func UniqEmail(obj interface{}) bool {
	var v *validator.Validate
	var topStruct reflect.Value
	var currentStructOrField reflect.Value
	var field reflect.Value
	var fieldType reflect.Type
	var fieldKind reflect.Kind
	var param string

	if email, ok := field.Interface().(string); ok {
		configs.DB.First(&obj, map[string]interface{}{"email": email})
		if obj.ID > 0 {
			return false
		}
	}
	return true
}
