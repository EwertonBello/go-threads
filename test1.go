package main
import (
    "fmt"
    "time"
)
func hello(text string) {
    for i := 0; i < 5; i++ {
        time.Sleep(200 * time.Millisecond)
        fmt.Println(i+1, text)
    }
}

func fibonacci() func() int {
    first, second := 0, 1
    return func() int {
        ret := first
        first, second = second, first+second
        return ret
    }
}

func th_fibo(resp chan string) {
	f := fibonacci()
	for i := 0; i < 100000000; i++ {
		fmt.Println(i+1,"-",f())
	}

	resp <- "Fim da fibo"
}

func main() {
    go hello("world")
    // hello("hello")
    time.Sleep(2000 * time.Millisecond)
    fmt.Println("Final")
    resp := make(chan string)
    go th_fibo(resp)
    fmt.Println(<-resp)
}
