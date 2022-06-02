package main

import(
    "fmt"
)

type Int int

func (x Int) plus(y Int) Int{
    return x+y
}

func main(){
    var x Int=10
    var y Int=2
    fmt.Println(x)
    fmt.Println(x.plus(y))
}
