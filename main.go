package main

import (
    "os"
    "log"
    "fmt"
    "io/ioutil"
)

func main() {
    // Open file for reading
    file1, err := os.Open("files/text1.txt")
    file2, err := os.Open("files/text2.txt")
    if err != nil {
        log.Fatal(err)
    }

    // os.File.Read(), io.ReadFull(), and
    // io.ReadAtLeast() all work with a fixed
    // byte slice that you make before you read

    // ioutil.ReadAll() will read every byte
    // from the reader (in this case a file),
    // and return a slice of unknown slice
    data1, err := ioutil.ReadAll(file1)
    data2, err := ioutil.ReadAll(file2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("1")
    fmt.Println("Data as Unicode: %U\n", data1)
    fmt.Printf("Data as hex: %x\n", data1)
    fmt.Printf("Data as string: %s\n", data1)
    fmt.Println("Number of bytes read:", len(data1))
    fmt.Println("")
    fmt.Println("")
    fmt.Println("Data as Unicode: %U\n", data2)
    fmt.Printf("Data as hex: %x\n", data2)
    fmt.Printf("Data as string: %s\n", data2)
    fmt.Println("Number of bytes read:", len(data2))


}
