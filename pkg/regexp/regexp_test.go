package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCreateRandomNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		r := showlistCreateValueBy("random_integer(5,10)")
		fmt.Println(r)
	}
}

var paraRegex *regexp.Regexp

func TestMatch1(t *testing.T) {
	paraRegex, _ = regexp.Compile("{{(.*?)}}")
	fmt.Println(getConditionsOfField("{{user_name.length<10}}...{{user_name_en.length>10}}"))

	findMatch("user_name.length<10")
	findMatch("user_name.length=10")
	findMatch("user_name.length>10")
}

func findMatch(text string) {
	reg, _ := regexp.Compile(`(.*?)\.(.*?)([><=])(\d+)`)
	arr := reg.FindAllStringSubmatch(text, -1)
	for _, it := range arr {
		fmt.Println(it)
	}
}

func getConditionsOfField(expression string) []condition {
	arr := paraRegex.FindAllStringSubmatch(expression, -1)
	if len(arr) == 0 {
		return nil
	}

	reg, _ := regexp.Compile(`(.*?)\.(.*?)([><=])(\d+)`)

	var conditions []condition
	for _, it := range arr {
		if len(it) < 1 {
			continue
		}
		conditionArr := reg.FindAllStringSubmatch(it[1], -1)
		if len(conditionArr) > 0 && len(conditionArr[0]) > 4 {
			conditions = append(conditions, condition{
				field:    conditionArr[0][1],
				property: conditionArr[0][2],
				operator: conditionArr[0][3],
				value:    conditionArr[0][4],
			})
		}
	}
	return conditions
}

// condition user_name length < 10
type condition struct {
	field    string
	property string
	operator string
	value    string
}
