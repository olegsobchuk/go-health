package validators

import (
	"reflect"

	"github.com/olegsobchuk/go-health/configs"
)

// UniqEmail add unique validation to field *email*
func UniqEmail(obj interface{}) bool {
	var field reflect.Value

	if email, ok := field.Interface().(string); ok {
		configs.DB.First(&obj, map[string]interface{}{"email": email})
		if obj.(struct{ ID uint }).ID > 0 {
			return false
		}
	}
	return true
}
