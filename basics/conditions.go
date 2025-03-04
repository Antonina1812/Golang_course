package basics

import "fmt"

func ConditionsPrint() {
	// блок инициализации перед условием
	m := map[string]string{"name": "Tonya"}
	if keyValue, keyExists := m["name"]; keyExists {
		fmt.Println("name =", keyValue)
	}

	val := "name"
	switch val {
	case "name":
		fallthrough // проваливание в следующее условие
	case "date":
		fmt.Println(val)
	default:
		fmt.Println("default")
	}

	val1, val2 := 1, 2
	switch {
	case val1 > 0 && val2 < 4:
		fmt.Println("good")
	case val1 > 100:
		fmt.Println("not good")
	}

	//Loop:
	for key, val := range m {
		fmt.Println("switch in loop", key, val)
	Loop: //добавляем loop, если захотим выйти из цикла или switch'а
		switch {
		case key == "lastname":
			break Loop
		case key == "name":
			fmt.Println(key)
			break Loop
		}
	}
}
