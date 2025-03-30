package json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Phone    string
}

var jsonStr = `{"id": 1, "username": "tonya", "phone": "12345"}`

func JSONPrint() {
	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u) // из json в структуру
	fmt.Printf("struct:\n\t%#v\n\n", u)

	u.Phone = "999"
	result, _ := json.Marshal(u) // из структуры в json
	fmt.Printf("json string:\n\t%s\n", string(result))
}
