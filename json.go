package main

import(
	"encoding/json"
	"fmt"
)

type uniono struct{
	a int
	b string
	c [3]string
	d map[string]string
}

func func main() {
	t:=uniono(
		a=1,
		b="b",
		c=[3]string{"a","b","c"},
		d=map[string]string{""}
		)
}