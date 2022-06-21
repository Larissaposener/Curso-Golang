//hello.go
package main

import "fmt"

func main() {
 



    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scanf("%d",&comando)
  

    switch comando{
        case 1:
        fmt.Println("Monitorando...")
        case 2:
        fmt.Println("Exibindo Logs...")
        case 3:
        fmt.Println("Saindo do programa...")
        default:
        fmt.Println("nenhuma das opções")
    }
}