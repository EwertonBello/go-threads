package main
import (
    "fmt"

)
func hello(name string, text string, final chan string) {
    // for i := 0; i < 100; i++ {
    //     time.Sleep(10000 * time.Millisecond)
    //     fmt.Println(i+1, text)
    // }
    for i := 0; i < 100000000; i++ {
        fmt.Println(i+1, name)
    }
    final <- text+" - Finalizado"
}
func main() {
    final := make(chan string)
    go hello("Thread 1", "Hello", final)
    go hello("Thread 2", "World", final)
    for i := 0; i < 100000000; i++ {
        fmt.Println(i+1, "Principal")
    }
    fmt.Println(<-final)
    fmt.Println(<-final)
}
