package main

import (
  "fmt"
  "io/ioutil"
  "bytes"
)

func main() {

        file1,_:= ioutil.ReadFile("file/file1.txt")
        file2,_:= ioutil.ReadFile("file/file2.txt")
        byteslice1 := []byte(file1)
        byteslice2 := []byte(file2)


  fmt.Println(bytes.Equal(byteslice1, byteslice2))
}
