package util

import (
	"reflect"
)

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}