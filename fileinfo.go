package main

import (
  "fmt"
  "log"
  "os"
)



func main() {
  arg1 := os.Args[1]
  FullFileInfo(arg1)

}

func FullFileInfo(filename string) {

  var (
    fileInfo os.FileInfo
    err      error
  )
  // stat returnerer file info.
  // error om det ikke er noe fil.
  fileInfo, err = os.Stat(filename)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("File name:", fileInfo.Name())

}
