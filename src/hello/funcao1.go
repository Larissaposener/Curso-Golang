//hello.go
package main

import "fmt"
import "os"

func main() {
    //chamndo a função
    exibeIntroducao()
    exibeMenu()
 
   var comando int
    comando = leComando();
    switch comando{
        case 1:
        fmt.Println("Monitorando...")
        case 2:
        fmt.Println("Exibindo Logs...")
        case 0:
        fmt.Println("Saindo do programa...")
        os.Exit(0)
        default:
        fmt.Println("nenhuma das opções")
        os.Exit(-1)
    }
}

func exibeIntroducao(){
    var nome string = "Larissa"
    var idade int = 19
    var versao float32 = 1.1
    fmt.Println("hello, meu nome é " + nome +" e curso analise e desenvolvimento de sitema")
    fmt.Println("hello, minha versão é", versao)
    fmt.Println("hello, minha idade é", idade)
    
}
func exibeMenu(){
    
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
    
}
//int informa que a função vai retornar um int
func leComando() int{
      var comandoLido int
      fmt.Scanf("%d",&comandoLido)
      fmt.Println("comando recebe", comandoLido)
      
      return comandoLido
    
}