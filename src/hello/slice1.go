//hello.go
package main

import "fmt"
import "os"
//import "net/http"

func main() {
    exibeNome()

}

func iniciarMonitoramento(){
    fmt.Println("monitorando...")
    var sites [4]string 
    sites[0] = "https://random-status-code.herokuapp.com/"
    sites[1] = "https://www.alura.com.br"
    sites[2] = "https://www.caelum.com.br"
    
    fmt.Println(sites)

   // site := "https://cursos.alura.com.br"
   // resp, _ := http.Get(site)
   // fmt.Println(resp)
    
    // if resp.StatusCode == 200 {
      //  fmt.Println("Site:", site, "foi carregado com sucesso!")
    //} else {
        //fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
   // }
    
    
}

func exibeNome(){
    nomes := []string {"Larissa","Leticia","Lukayan","Luara"}
    fmt.Println(nomes)
   
    fmt.Println("o slice tem ",len(nomes), "itens com a capacidade de suportar", cap(nomes), "itens\n")
   
    
    nome := append(nomes,"Gildete","Sergio")
    fmt.Println(nome)
    fmt.Println("o slice tem ",len(nome), "itens com a capacidade de suportar", cap(nome), "itens")
}