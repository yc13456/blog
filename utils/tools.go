package utils

import (
	"errors"
	"fmt"
	"reflect"
	"unicode"
)

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func ChineseCheck(data any) error{
	v :=reflect.ValueOf(data)
	fmt.Println("kind: ",v.Kind())
	if v.Kind() != reflect.Struct{
		return errors.New("data is not struct")
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println("str: ",f.String())
			if IsChinese(f.String()){
				return errors.New(fmt.Sprintf("Chinese characters found, %s",f.String()))
			}
		}
	}
	return nil
}