package converters

import (
	"fmt"
	"regexp"
	"testing"
)

var UrlData = make(map[string][]string)
var pattern string

var UrlString string

var key string
var value string

func init() {

	key = "key"
	value = "value"

	for i := 1; i < 2; i++ {

		_key := fmt.Sprintf("%s%d", key, i)

		_value1 := fmt.Sprintf("%s%d", value, i)
		_value2 := fmt.Sprintf("%s%d", value, i+1)

		UrlData[_key] = append(UrlData[_key], _value1, _value2)
	}

	pattern = fmt.Sprintf(`{"%s1":\["%s1","%s2"]}`, key, value, value)
}

func TestUrlValuesToString(t *testing.T) {

	str := UrlValuesToString(UrlData)

	want := regexp.MustCompile(pattern)

	if !want.MatchString(str) {
		t.Fatalf("Could not find the pattern in the string str & pattern: %s %s", str, pattern)
	}
}

func TestStringToUrlValues(t *testing.T) {

	str := UrlValuesToString(UrlData)
	data := StringToUrlValues(str)

	want := fmt.Sprintf("%s%d", value, 1)
	got := data[fmt.Sprintf("%s%d", key, 1)][0]

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
