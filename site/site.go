package main

import(
	"time"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main(){

	app:=iris.New()

	iris.RegisterOnInterrupt(func() {
	    timeout := 5 * time.Second
	    ctx, cancel := context.WithTimeout(context.Background(), timeout)
	    defer cancel()
	    // close all hosts
	    app.Shutdown(ctx)
	})

	customLogger := logger.New(logger.Config{
	Status: true,
	IP: true,
	Method: true,
	Path: true,
	Query: true,

	MessageContextKeys: []string{"logger_message"},

	//MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)

	app.Logger().SetLevel("debug")

	app.HandleDir("/",".")
	app.HandleDir("/external","./external")
	app.HandleDir("/icon","./icon")
	app.HandleDir("/pages","./pages")

	app.Run(iris.Addr(":8000"),iris.WithoutInterruptHandler)
}