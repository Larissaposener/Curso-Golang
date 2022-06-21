package main 
import "fmt"

func main() {
  nome,idade := devolveNomeEIdade()
  fmt.Println(nome,idade)
}
  
func devolveNomeEIdade() (string, int){
     nome := "larissa"
     idade := 19
     return nome,idade
}