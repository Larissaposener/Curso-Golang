package main

import "fmt"
import "reflect"// para saber o tipo da variavel

func main() {
    var nome  = "Larissa"
    var idade int = 19
    var versao float32 = 1.1
    fmt.Println("hello, meu nome é " + nome +" e curso analise e desenvolvimento de sitema")
    fmt.Println("hello, minha versão é", versao)
    fmt.Println("hello, minha idade é", idade)
    //reflect.TypeOf saber o tipo da variavel
     fmt.Println("o tipo da varivel nome é", reflect.TypeOf(nome))
}