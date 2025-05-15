package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "os/signal"
)


func main() {
    var signalChan chan os.Signal = make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt)
    go func() {
        <-signalChan
        fmt.Println("\nleaving...")
        os.Exit(0)
    }()

    var reader *bufio.Reader = bufio.NewReader(os.Stdin)
    var input string = ""

    for input != "exit" || input != "quit" {
        fmt.Print("> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("error: ", err)
            os.Exit(1)
        }
        fmt.Println(strings.TrimSpace(input))
    }
}
