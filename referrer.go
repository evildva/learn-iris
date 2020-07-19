package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    app.Get("/", func(ctx iris.Context) {
        r := ctx.GetReferrer()
        switch r.Type {
        case iris.ReferrerSearch:
            ctx.Writef("Search %s: %s\n", r.Label, r.Query)
            ctx.Writef("Google: %s\n", r.GoogleType)
        case iris.ReferrerSocial:
            ctx.Writef("Social %s\n", r.Label)
        case iris.ReferrerIndirect:
            ctx.Writef("Indirect: %s\n", r.URL)
        }
    })

    app.Listen(":8080")
}