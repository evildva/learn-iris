package main
import(
  "fmt"
  "reflect"
  "net/http"
  "io/ioutil"
  "unsafe"
  //"compress/gzip"
)

func main(){

  client:=&http.Client{

  }
/*
  response,err:=http.Get("https://golang.google.cn/pkg/")
  if err!=nil {
    fmt.Println("Get error :%V",err)
    return
  }
  b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
  fmt.Println(*(*string)(unsafe.Pointer(&b)))

  fmt.Println(reflect.TypeOf(response))
  fmt.Println(response.Body)
  fmt.Println(response.ContentLength)
  */
  /*
  for k,v:=range client.Header{
    fmt.Println(k,"  ",v)
  }
  */
  //defer response.Body.Close()

  request,err:=http.NewRequest("GET","http://localhost:8000",nil)
  request.Header.Add("accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,/*;q=0.8")
  //request.Header.Add("accept-encoding","gzip, deflate, br")
  request.Header.Add("accept-language","zh-CN,zh;q=0.9,en;q=0.8,la;q=0.7,ko;q=0.6,ja;q=0.5")
  request.Header.Add("user-agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
  resp,err:=client.Do(request)
  fmt.Println(resp.Body)
  fmt.Println("response ",reflect.TypeOf(resp.Body))

  a, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
  fmt.Println("ioutil ",reflect.TypeOf(a))
  fmt.Println(*(*string)(unsafe.Pointer(&a)))

  //fmt.Println(a)

  defer resp.Body.Close()
}
