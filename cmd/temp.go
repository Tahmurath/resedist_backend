package cmd

//
//import (
//	"fmt"
//	"regexp"
//	"resedist/pkg/converters"
//)
//
//package main
//
//import (
//	"fmt"
//	"regexp"
//	"resedist/pkg/converters"
//)
//
//var UrlData = make(map[string][]string)
//
//func main() {
//	//UrlData = map[string][]string{}
//
//	key := "title"
//	key2 := "title2"
//	value := "body"
//	value2 := "body2"
//
//	pattern := `{"title":\["body`
//
//	UrlData[key] = []string{value, value2}
//
//	UrlData[key2] = append(UrlData[key2], value, value2)
//
//	fmt.Println(UrlData)
//
//	str := converters.UrlValuesToString(UrlData)
//	fmt.Println(str)
//
//	want := regexp.MustCompile(pattern)
//	if !want.MatchString(str) {
//		t.Fatalf(str)
//	}
//}
