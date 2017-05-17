package main

import (
    "os"
    "log"
    "fmt"
    "io/ioutil"
)

func main() {
    //Åpner filen
    file, err := os.Open("pg100.txt")
    if err != nil {
        log.Fatal(err)
    }


    // ioutil.ReadAll() Leser hele filen og printer ut data om bytes lest.
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Data as hex: %x\n", data)
    fmt.Printf("Data as string: %s\n", data)
    fmt.Println("Number of bytes read:", len(data))
}
