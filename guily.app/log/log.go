package log

import(
  "strings"
  "time"
  "os"
)

func todayFilename() string {
    nowtime := time.Now().Format("2006-01-02-15:04:05")
    return strings.Replace(nowtime,":","-",-1) + ".txt"
}

func LogFile(d string) *os.File {

    dir,_:=os.Getwd()//获得当前路径

    today:=time.Now().Format("2006-01-02")
    filename := todayFilename()

    dir=dir+d+today+"/"
    _, err := os.Stat(dir)//文件夹是否存在
    if err==nil{
        goto createfile
    }

    err = os.Mkdir(dir, os.ModePerm)//不存在就创建
    if err!=nil{
        panic(err)
    }
    // Open the file, this will append to the today's file if server restarted.
createfile:
    f, err := os.OpenFile(dir + filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }

    return f
}
