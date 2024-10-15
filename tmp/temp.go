package main

import (
	"fmt"
	"resedist/pkg/converters"
)

func hasKeyValue(m map[string][]string, key string, value string) bool {
	if v, exists := m[key]; exists {
		for _, val := range v {
			if val == value {
				return true
			}
		}
	}
	return false
}

func main() {
	var UrlData = make(map[string][]string)

	key := "title"
	key2 := "title2"
	value := "body"
	value2 := "body2"

	UrlData = map[string][]string{
		key:  {value, value2},
		key2: {value, value2},
	}

	fmt.Println(UrlData)

	str := converters.UrlValuesToString(UrlData)
	fmt.Println(str)

	UrlData2 := converters.StringToUrlValues(str)
	fmt.Println(UrlData2)

	str = converters.UrlValuesToString(UrlData2)
	fmt.Println(str)

}
