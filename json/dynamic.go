package json

import (
	"encoding/json"
	"fmt"
)

var JsonStr = `[
	{"id": 1, "username": "tonechka", "phone": 666},
	{"id": 2, "address": "none", "company": "google"}
]`

func DynamicPrint() {
	data := []byte(JsonStr)

	var user1 interface{}
	json.Unmarshal(data, &user1) // распаковка
	fmt.Printf("unpacked in empty interface:\n%#v\n\n", user1)

	user2 := map[string]interface{}{
		"id":       42,
		"username": "enot",
	}

	var user2i interface{} = user2
	result, _ := json.Marshal(user2i) // запаковка
	fmt.Printf("json string from map:\n%s\n", string(result))
}
