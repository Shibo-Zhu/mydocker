package main

import (
    "fmt"
    "time"
)

func worker(id int, ch chan string) {
    time.Sleep(time.Second)
    ch <- fmt.Sprintf("Worker %d done", id)
}

func main() {
    ch := make(chan string)

    for i := 1; i <= 5; i++ {
        go worker(i, ch)
    }

    for i := 1; i <= 5; i++ {
        fmt.Println(<-ch)
    }
}
