package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"unicode"
)

func InArray(val any, array any) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				exists = true
				return
			}
		}
	}
	return
}

func PascalToSnake(pascal string) string {
	var words []string
	l := 0
	for s := pascal; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, s[:l])
	}
	snake := strings.ToLower(strings.Join(words, "_"))
	return snake
}

func PHPInArray(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func InterfaceToStruct[R any](params any) (R, error) {
	reqBodyBytes := new(bytes.Buffer)
	var result R
	if err := json.NewEncoder(reqBodyBytes).Encode(params); err != nil {
		return result, err
	}
	if err := json.Unmarshal(reqBodyBytes.Bytes(), &result); err != nil {
		return result, err
	}
	return result, nil
}
