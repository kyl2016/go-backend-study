package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	// Create a string using ` string escape ticks
	// query := `{"$eq":"last value"}`
	query := `{"$or": [ "c1", {"$and": ["c3", {"$or": ["c2", "c4", "c5"]}]}]}`

	values := map[string]bool{
		"c1": false,
		"c2": true,
		"c3": true,
		"c4": false,
		"c5": false,
	}

	// Declare an empty BSON Map object
	var bsonMap bson.M

	// Use the JSON package's Unmarshal() method
	err := json.Unmarshal([]byte(query), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	} else {
		stack := parse(0, map[string]interface{}(bsonMap))
		result := values[stack[len(stack)-1].item.(string)]

		i := len(stack) - 2
		for i >= 0 {
			condition := stack[i]
			v := values[condition.item.(string)]
			fmt.Println(condition.operator, v)

			switch condition.operator {
			case "$or":
				result = result || v
			case "$and":
				result = result && v
			}
			i--
		}
		fmt.Println(result)

		// for k, v := range bsonMap {
		// 	fmt.Println("bsonMap:", bsonMap)
		// 	fmt.Println("bsonMap TYPE:", reflect.TypeOf(bsonMap))
		// 	fmt.Println("BSON:", reflect.TypeOf(bson.M{"int field": bson.M{"$gt": 42}}))
		// }
	}
}

func parse(level int, data map[string]interface{}) []condition {
	var r []condition
	for k, v := range data {
		if arr, ok := v.([]interface{}); ok {
			for _, v := range arr {
				if m, ok := v.(map[string]interface{}); ok {
					level++
					r = append(r, parse(level, m)...)
				} else {
					r = append(r, condition{k, v, level})
				}
			}
		} else {
			r = append(r, condition{k, v, level})
		}
	}
	return r
}

type condition struct {
	operator string
	item     interface{}
	level    int
}
