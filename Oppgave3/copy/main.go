package main

import (
    "os"
    "log"
    "io"
    "fmt"
    "bufio"

)

// Kopierer filen og scanner etter antall linjer i teksten
func main() {

    originalFile, err := os.Open("pg100.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer originalFile.Close()

    // Oppretter en ny fil
    newFile, err := os.Create("pg100.txt_copy")
    if err != nil {
        log.Fatal(err)
    }
    defer newFile.Close()



    // Kopierer bytes fra orginal filen til den nye filen
    bytesWritten, err := io.Copy(newFile, originalFile)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Copied %d bytes.", bytesWritten)

    // scanner antall linjer

    file, _ := os.Open("pg100.txt_copy")
fileScanner := bufio.NewScanner(file)
lineCount := 0
for fileScanner.Scan() {
    lineCount++
}
fmt.Println("number of lines:", lineCount)




    // Commit the file contents
    // Flushes memory to disk
    err = newFile.Sync()
    if err != nil {
        log.Fatal(err)
    }
  }
