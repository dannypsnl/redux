package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `valid:"[0-9]+"`
	Name string `valid:"required"`
}

func main() {
	user := User{
		Id:   1,
		Name: "Jason Json",
	}
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("valid")

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
