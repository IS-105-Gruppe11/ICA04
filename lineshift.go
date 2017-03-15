package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
  arg1 := os.Args[1]
  IsMacOrWin(arg1)

}

func IsMacOrWin(filename string) {
	file1, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}


	data1, err := ioutil.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
	}


	match1, err := regexp.Match(`\r\n`, data1)
	if err != nil {
		log.Fatal(err)
	}
	text1Platform := "Windows"
	if !match1 {
		text1Platform = "Mac"
	}


	fmt.Printf("Text1 is using %s line breaks\n", text1Platform)


}
