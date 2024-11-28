package main

import (
	"fmt"
	"reflect"
)

func defineType(i interface{}) string {
	switch i.(type) {
	// if we need value, we can get it by switch v := i.(type)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return "int"
	case string:
		return "string"
	case bool:
		return "boolean"
	case chan interface{}:
		return "chanel interface{}"
	default:
		var isChannel = reflect.ValueOf(i).Kind() == reflect.Chan
		// typeByChannel := reflect.ValueOf(i).Type().Elem()
		typeByChannel := reflect.ValueOf(i).Type()
		if isChannel {
			return fmt.Sprintf("chanel %s", typeByChannel)
		} else {
			return typeByChannel.String()
		}
	}
}
