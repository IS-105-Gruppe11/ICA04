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
  IsLineshiftMacOrWin(arg1)

}

func IsLineshiftMacOrWin(filename string) {
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


	fmt.Printf("The textfile is using %s line breaks\n", text1Platform)


}
