package main

import(//iris doc https://godoc.org/github.com/kataras/iris#Configuration
  //"fmt"
  //"os"
  "time"
  "context"
  "github.com/kataras/iris/v12"
  //"github.com/kataras/golog"
	//"github.com/kataras/iris/v12/middleware/logger"
  //"github.com/google/wire/cmd/wire"
  "guily.app/log"
)

func main(){
  app:=iris.New()

  logfile:=log.LogFile("/log/")//设置log输出目录为 PWD/log/
  defer logfile.Close()

  iris.RegisterOnInterrupt(func() {
        timeout := 5 * time.Second
        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()
        // close all hosts
        app.Shutdown(ctx)
    })
  app.Logger().SetLevel("debug")
  app.Logger().SetOutput(logfile)

  //golog.New().SetOutput(os.Stdout)

  app.Get("/",func (ctx iris.Context)  {
    ctx.WriteString("Server"/*log.TodayFilename()*/)
  })

  app.Get("/download", func(ctx iris.Context) {
		file := "main.go"
		ctx.SendFile(file, "main.go")
	})

  app.Get("/test",func(ctx iris.Context){
    ctx.WriteString("abc  test")
  })

  app.Run(iris.Addr(":8000"),iris.WithoutInterruptHandler)
}
