package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "log"
)

func main() {
    arg1 := os.Args[1:]
    if len(arg1) == 1 {
      IsLineshift(arg1[0])
    }
}

func IsLineshift(filename string)  {

 winLineshift := false

file1, err := os.Open(filename)
if err != nil {
    log.Fatal(err)
  }
data1, err := ioutil.ReadAll(file1)
if err != nil {
    log.Fatal(err)
  }

var c byte = 13

for i := 0; i < len(data1) || winLineshift == true; i++ {
  if data1[i] == c {
    winLineshift = true
  } else {
    winLineshift = false
  }

}
if winLineshift == true {
  fmt.Println("Windows")
} else {
  fmt.Println("Mac")
}
}
