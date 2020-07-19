package main

import(
  "fmt"
  "github.com/kataras/iris/v12"
)

func main(){
	app:=iris.New()

	app.Get("/",func(ctx iris.Context){
		fmt.Println(ctx.Host)
		fmt.Println(ctx.Path())
		fmt.Println(ctx.RemoteAddr)
		fmt.Println(ctx.FullRequestURI)
		fmt.Println(ctx.URLParams())
		s,_:=ctx.GetBody()
		fmt.Println(string(s[:]))
		})

	app.Post("/",func(ctx iris.Context){
		fmt.Println(ctx.Host)
		fmt.Println(ctx.Path())
		fmt.Println(ctx.RemoteAddr)
		fmt.Println(ctx.FullRequestURI)
		fmt.Println(ctx.URLParams())
		s,_:=ctx.GetBody()
		fmt.Println(string(s[:]))
		})

	app.Run(iris.Addr(":8800"))
}