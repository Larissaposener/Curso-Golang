//hello.go
package main

import "fmt"

func main() {

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scanf("%d",&comando)
   
    if comando == 0{
        fmt.Println("Monitorando...")
    }else if comando == 1{
        fmt.Println("valor digitado é 1")
    }else if comando == 2{
        fmt.Println("valor digitado é 2")
    }
}