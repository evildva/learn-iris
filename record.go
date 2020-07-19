package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    // start record.
    app.Use(func(ctx iris.Context) {
        ctx.Record()
        ctx.Next()
    })

    app.Get("/", func(ctx iris.Context) {
    ctx.WriteString("success")
    ctx.Next() // calls the Done middleware(s).
    })

    // collect and "log".
    app.Done(func(ctx iris.Context) {
        body := ctx.Recorder().Body()

        // Should print success.
        app.Logger().Infof("sent: %s", string(body))
    })

    app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}