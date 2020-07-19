package main

import(
"fmt"
)

type callback func(s string)

func print(f callback,s string){
  f(s)
}

func p(s string){
  fmt.Println(s)
}

func main(){

  print(p,"callback")

}
