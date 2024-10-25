package utils

import (
	"fmt"
	"reflect"
)

func StructToInterfaceObj(structWithData interface{}, removeColumn []string) map[string]interface{} {
	// new interface layout would be a place
	// for injecting new data and value of the data
	newInterfaceLayout := make(map[string]interface{})

	// Get the datatype of struct
	structDataType := reflect.TypeOf(structWithData)
	// Get the datatype of struct
	structDataValue := reflect.ValueOf(structWithData)

	// Loop by counting all column
	for i := 0; i < structDataType.NumField(); i++ {
		getFieldName := structDataType.Field(i).Name
		dataFieldTag := structDataType.Field(i).Tag
		// this "InArray" function is on this link : https://codefile.io/f/8LN6E9KBe6
		if PHPInArray(removeColumn, getFieldName) < 0 {
			fieldValue := structDataValue.Field(i)
			fieldValueToString := fmt.Sprintf("%v", fieldValue)
			newInterfaceLayout[dataFieldTag.Get("json")] = fieldValueToString
		}
	}

	return newInterfaceLayout
}
