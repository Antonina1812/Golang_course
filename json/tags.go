package json

import (
	"encoding/json"
	"fmt"
)

type NewUser struct {
	ID       int `json:"user_id,string"` // метаинформация (тэг)
	Username string
	Address  string `json:",omitempty"` // "omitempty" - если поле пустое, то его не нужно писать
	Company  string `json:"-"`          // "-" - поле вообще не нужно трогать
}

func TagsPrint() {
	u := &NewUser{
		ID:       42,
		Username: "tonya",
		Address:  "",
		Company:  "google",
	}
	result, _ := json.Marshal(u)
	fmt.Printf("json string:\n\t%s\n", string(result))
	// {"user_id":"42","Username":"tonya"}
}
