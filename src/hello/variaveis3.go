package main

import "fmt"
import "reflect"// para saber o tipo da variavel

func main() {
    nome  := "Larissa"//declara e informa o tipo

    fmt.Println("hello, meu nome é ", nome )
  
    //reflect.TypeOf saber o tipo da variavel
     fmt.Println("o tipo da varivel nome é", reflect.TypeOf(nome))
}