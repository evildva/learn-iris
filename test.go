/*
package main

import (
        "encoding/json"
        "fmt"
        "os"
)

func main() {
        type ColorGroup struct {
                ID     int
                Name   string
                Colors []string
        }
        
        type picture struct{
                aA       string `json:"a"`
                uUrls   []string `json:"b"`
        }
        
        group := ColorGroup{
                ID:     1,
                Name:   "Reds",
                Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
        }
        
        js:=picture{
        aA:"abc",
        uUrls: []string{"aa","bb","cc"},
        }
        
        b, err := json.Marshal(group)
        if err != nil {
                fmt.Println("error:", err)
        }
        
        c, err := json.Marshal(js)
        if err != nil {
                fmt.Println("error:", err)
        }
        
        os.Stdout.Write(b)
        
        os.Stdout.Write(c)
}

package main

import (
        "github.com/kataras/iris/v12"
        "github.com/kataras/iris/v12/core/router"
)

func main() {
        app := iris.New()
        // need for manually reverse routing when needed outside of view engine.
        // you normally don't need it because of the {{ urlpath "routename" "path" "values" "here"}}
        rv := router.NewRoutePathReverser(app)

        myroute := app.Get("/anything/{anythingparameter:path}", func(ctx iris.Context) {
                paramValue := ctx.Params().Get("anythingparameter")
                ctx.Writef("The path after /anything is: %s", paramValue)
        })

        myroute.Name = "myroute"

        // useful for links, although iris' view engine has the {{ urlpath "routename" "path values"}} already.
        app.Get("/reverse_myroute", func(ctx iris.Context) {
                myrouteRequestPath := rv.Path(myroute.Name, "any/path")
                ctx.HTML("Should be <b>/anything/any/path</b>: " + myrouteRequestPath)
        })

        // execute a route, similar to redirect but without redirect :)
        app.Get("/execute_myroute", func(ctx iris.Context) {
                ctx.Exec("GET", "/anything/any/path") // like it was called by the client.
        })

        // http://localhost:8080/reverse_myroute
        // http://localhost:8080/execute_myroute
        // http://localhost:8080/anything/any/path/here
        //
        // See view/template_html_4 example for more reverse routing examples
        // using the reverse router component and the {{url}} and {{urlpath}} template functions.
        app.Run(iris.Addr(":8080"))
}

package main

import(
        "fmt"
        "go.mongodb.org/mongo-driver/bson"
)

type t struct{
        aa string                          `bson:"aa"`
}

type td struct{
        a string                           `bson:"a"`
        b int                              `bson:"b"`
        c map[string]interface{}           `bson:"c"`
        d []string                         `bson:"d"`
        e t                                `bson:"e"`
}

func main(){
        //d:=bson.D{{"a","abc"},{"b",111},{"d",[]string{"a","b","c"}},{"e","aaa"}}
        ntd:=td{a:"ntd",b:222,c:map[string]interface{}{"a":1,"b":2},d:[]string{"d","e"},e:t{"tt"}}
        fmt.Printf("%v",ntd)
        a,_:=bson.Marshal(ntd)
        _,b,_:=bson.MarshalValue(ntd)
        fmt.Printf("%v",string(a))
        fmt.Printf("%v",b)
        var m td
        bson.Unmarshal(a,&m)
        fmt.Println(m)
}
*/
package main

import(
        "os"
        "github.com/kataras/iris/v12"
)

func main(){
        app:=iris.New()
        app.Get("/",func (ctx iris.Context){
        str,_:=ReadAll("./assets/annoce.html")
        ctx.HTML(str)
    })
        app.Run(iris.Addr(":8000"))
}