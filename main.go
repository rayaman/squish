package main

import "github.com/rayaman/squish/encoding"

type User struct {
	Name string `binary:"limit=24"` // no
	Age  int    `binary:""`
}

func main() {
	user := &User{Name: "Ryan Ward", Age: 28}
	encoding.GetOptions(user)
}
