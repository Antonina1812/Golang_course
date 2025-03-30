package json

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

// reflect позволяет перебирать поля структуры в цикле и работать с данными этих полей

type User1 struct {
	ID       int
	RealName string `unpack:"-"` // тэг значит, что поле не нужно использовать для распаковки
	Login    string
	Flags    int
}

func Reflect(u interface{}) error {
	val := reflect.ValueOf(u).Elem() // сама структура

	fmt.Printf("%T have %d field:\n", u, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Printf("\tname=%v, type=%v, value=%v, tag=%v\n",
			typeField.Name,
			typeField.Type.Kind(),
			valueField,
			typeField.Tag,
		)
	}
	return nil
}

func UnpackReflect(u interface{}, data []byte) error {
	r := bytes.NewReader(data)
	val := reflect.ValueOf(u).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Tag.Get("unpack") == "-" {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			var value uint32
			binary.Read(r, binary.LittleEndian, &value) // чтение данных в переменную
			valueField.Set(reflect.ValueOf(int(value))) // установка значения
		case reflect.String:
			var lenRaw uint32
			binary.Read(r, binary.LittleEndian, &lenRaw)

			dataRaw := make([]byte, lenRaw)
			binary.Read(r, binary.LittleEndian, &dataRaw)

			valueField.SetString(string(dataRaw))
		default:
			return fmt.Errorf("bad type: %v for field %v", typeField.Type.Kind(), typeField.Name)
		}
	}
	return nil
}

func ReflectPrint() {
	u := &User1{
		ID:       1,
		RealName: "tonya",
		Flags:    32,
	}

	data := []byte{
		128, 36, 17, 0,

		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,

		16, 0, 0, 0,
	}

	// err := Reflect(u)
	// if err != nil {
	// 	panic(err)
	// }

	err := UnpackReflect(u, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", u)
}
