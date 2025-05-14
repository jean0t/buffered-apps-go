package main

import (
    "fmt"
    "bufio"
    "os"
)

func readContents(file *os.File) []string {
    var scanner *bufio.Scanner = bufio.NewScanner(file)
    var content []string = []string{}
    for scanner.Scan() {
       content = append(content, scanner.Text()) 
    }

    return content
}

func copyContents(destiny string, content []string) error {
    file, err := os.OpenFile(destiny, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    var writer *bufio.Writer = bufio.NewWriter(file)
    defer writer.Flush()

    for _, line := range content {
        _, err = writer.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }

    return nil
}

func isDir(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        fmt.Println("An error has occurred when verifying the target file")
        return false
    }

    if info.IsDir() {
        return true
    }

    return false
}

func usage() {
    fmt.Println("USAGE: cpf <origin_file> <destination_file>")
}

func main() {
    var args []string = os.Args
    if (args[1] == "-h" || args[1] == "--help") {
        usage()
        return
    }
    
    var origin_file_path string = args[1]
    if isDir(origin_file_path) {
        fmt.Println(origin_file_path, "must be a file")
        return
    }
    origin_file, err := os.Open(origin_file_path)
    if err != nil {
        fmt.Println("Error opening the file ", origin_file_path)
        return
    }

    var destination_file = args[2]
    var contents []string = readContents(origin_file)
    copyContents(destination_file, contents)

    fmt.Println("File copied")

}
