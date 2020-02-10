package shapes

import (
	"fmt"
	"reflect"
	"strings"
)

// DynamicFuncCall call the correct calculator function depending on shape and calculator names
func DynamicFuncCall(shape, calculator string, sizes []float64) {
	fns := map[string]interface{}{
		"triangleArea":  TriangleArea,
		"rectangleArea": RectangleArea,
	}
	fnName := fmt.Sprintf("%s%s", shape, strings.Title(calculator))
	fnValue := reflect.ValueOf(fns[fnName])
	params := make([]reflect.Value, fnValue.Type().NumIn())

	for index, size := range sizes {
		params[index] = reflect.ValueOf(size)
	}

	fnValue.Call(params)
}
