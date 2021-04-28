package main
import (
    "fmt"
    "time"
)

// Pronta, executando, bloqueada, finalizada

func hello(name string, resp chan string) {
    fmt.Println(name, "- Pronta!")
    fmt.Println(name, "- Executando")
    for i := 0; i < 5; i++ {
        fmt.Println(i+1, name)
    }
    fmt.Println(name, "- Bloqueada por 5 segundos")
    
    time.Sleep(5000 * time.Millisecond)
    resp <- name+" - Finalizada"
}

func main() {
    resp := make(chan string)
    for i := 0; i < 5; i++ {
        go hello("Thread "+fmt.Sprint(i+1), resp)
    }
    fmt.Println("Final do bloco da Thread Principal")
    fmt.Println(<-resp)
}
