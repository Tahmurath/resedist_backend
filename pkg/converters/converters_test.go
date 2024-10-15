package converters

import (
	"fmt"
	"regexp"
	"testing"
)

var UrlData = make(map[string][]string)

type UrlDataSet struct {
	row []UrlDataRow
}
type UrlDataRow struct {
	key   string
	value string
}

var key = "title"
var key2 = "title2"
var value = "body"
var value2 = "body2"

func init() {
	UrlData = map[string][]string{
		key:  {value, value2},
		key2: {value, value2},
	}
}

func TestUrlValuesToString(t *testing.T) {
	//var UrlData = make(map[string][]string)

	pattern := `{"title2":\["body`

	UrlData = map[string][]string{
		key:  {value, value2},
		key2: {value, value2},
	}
	fmt.Println(UrlData)
	//UrlData[key] = []string{value, value2}
	//UrlData[key2] = append(UrlData[key2], value, value2)

	str := UrlValuesToString(UrlData)
	fmt.Println(str)

	want := regexp.MustCompile(pattern)
	if !want.MatchString(str) {
		t.Fatalf(str)
	}
}
