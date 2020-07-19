package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app:=iris.New()
	app.HandleDir("/release","./release")
	app.Run(iris.Addr(":9000"))
}