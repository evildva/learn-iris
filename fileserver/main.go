package main

import(
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	//"net/http"
	//"path/filepath"
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

func getCurrentDirFiles(ctx iris.Context,path string){

    files, _ := ioutil.ReadDir(path)
    for _, f := range files {
    	path=strings.Replace(path,"/packages","/proj",1)
    	h:="<a href='"+"."+path+"/"+f.Name()+"'>"+f.Name()+"</a></br>"
    	ctx.HTML(h)
    }
}

func doWithFileType(ctx iris.Context,path string){
	
	file, err := os.Lstat(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("do path  "+file.Name())
	switch mode := file.Mode(); {
	case mode.IsRegular():
		//ctx.SendFile(path,path)
		content, _ := ioutil.ReadFile(path)
		s:="<div>"+string(content)+"</div></br>"
		ctx.WriteString(path)
		ctx.HTML(s)
		fmt.Println("file")
	case mode.IsDir():
		getCurrentDirFiles(ctx,path)
		fmt.Println("dir")
	case mode&os.ModeSymlink != 0:
		real,err:=os.Readlink(path)
		if err != nil {
		fmt.Println(err)
		}
		fmt.Println(real)
		getCurrentDirFiles(ctx,path)
		fmt.Println("link")
	case mode&os.ModeNamedPipe != 0:
	}
}

func main(){
	app := iris.New()

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

	app.HandleDir("/", "./packages")

	pro:=app.Party("/proj")

	pro.Get("/",func(ctx iris.Context){
		doWithFileType(ctx,"./packages")
	})

	pro.Get("/*",func(ctx iris.Context){
		path:=strings.Replace("."+ctx.Path(),"proj","packages",1)
		fmt.Println("get path  "+path)
		doWithFileType(ctx,path)
	})

	app.Listen(":8080")
}