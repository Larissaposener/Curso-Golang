//hello.go
package main

import "fmt"
import "os"
//import "net/http"

func main() {
   
    //essa função tem o loop infinito e valida o status de uma requisão de site
    //chamndo a função
   exibeIntroducao()
 
   for{
	exibeMenu()
   }
   
       var comando int
        comando = leComando();
        switch comando{
            case 1:
            iniciarMonitoramento()
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
      var comandoLido int = 1
      fmt.Scanf("%d",&comandoLido)
      fmt.Println("comando recebe", comandoLido)
      
      
      return comandoLido
    
}

func iniciarMonitoramento(){
    fmt.Println("monitorando...")
    sites := []string {"https://random-status-code.herokuapp.com/","https://www.alura.com.br","https://www.caelum.com.br"}
    
    for i, site := range sites{
        fmt.Println("estou passando na posição", i ,"do meu slice e essa posição obtém o site", site)
    }
  

   // site := "https://cursos.alura.com.br"
   // resp, _ := http.Get(site)
   // fmt.Println(resp)
    
    // if resp.StatusCode == 200 {
      //  fmt.Println("Site:", site, "foi carregado com sucesso!")
    //} else {
        //fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
   // }
    
    
}

