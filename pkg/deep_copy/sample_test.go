package main

import (
	"encoding/json"
	"testing"

	"github.com/kyl2016/Play-With-Golang/utility"
	"github.com/mohae/deepcopy"
)

func TestMapDeepCopy(t *testing.T) {
	body := `{"b":"daaf","d2ewq":"sadfasdf","value":["a", "b"],"sdf":false,"2w3er":1000000000000, "sldfj":1.1}`
	var src map[string]interface{}
	err := json.Unmarshal([]byte(body), &src)
	utility.PanicIfNotNil(err)

	copied := deepcopy.Copy(src)

	for k, v := range copied.(map[string]interface{}) {
		switch v := v.(type) {
		case []interface{}:
			for index, it := range v {
				if it != src[k].([]interface{})[index] {
					t.Error(k)
				}
			}
		default:
			if v != src[k] {
				t.Error(k)
			}
		}
	}
}
