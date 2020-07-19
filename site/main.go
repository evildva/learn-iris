package main

import (
	"html/template"
	"io/ioutil"
	//"strings"
	"unsafe"
	"log"
	"os"
)

func str2bytes(s string) []byte {
    x := (*[2]uintptr)(unsafe.Pointer(&s))
    h := [3]uintptr{x[0], x[1], x[1]}
    return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func main(){

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	content, err := ioutil.ReadFile("index.html")
	check(err)

	t, err := template.New("webpage").Parse(bytes2str(content))
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems)
	check(err)
}
