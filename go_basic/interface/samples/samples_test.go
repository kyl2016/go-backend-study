package samples

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	// var i interface{}
	// i = []interface{}{}
	// fmt.Println(i)

	p := People{Name: "Bob"}
	m := map[string]People{"-20": p, "20-50": p}
	fmt.Println(getRanges(m))
}

type People struct {
	Name string
}

type Range struct {
	From int
	To   int
}

func getRanges(in interface{}) (map[string]Range, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	r := make(map[string]Range, len(m))
	for k := range m {
		r[k] = parseRange(k)
	}
	return r, nil
}

func parseRange(in string) Range {
	return Range{0, 20}
}
