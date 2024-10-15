package converters

import (
	"fmt"
	"regexp"
	"testing"
)

var UrlData = make(map[string][]string)

func TestUrlValuesToString(t *testing.T) {
	key := "title2"
	key2 := "title2"
	value := "body"
	value2 := "body2"

	pattern := `{"title2":\["body`

	UrlData[key] = []string{value, value2}

	UrlData[key2] = append(UrlData[key2], value, value2)

	fmt.Println(UrlData)

	str := UrlValuesToString(UrlData)
	fmt.Println(str)

	want := regexp.MustCompile(pattern)
	if !want.MatchString(str) {
		t.Fatalf(str)
	}
}
