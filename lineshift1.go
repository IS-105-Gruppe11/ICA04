package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	file1, err := os.Open("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}

	file2, err := os.Open("files/text2.txt")
	if err != nil {
		log.Fatal(err)
	}

	data1, err := ioutil.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
	}

	data2, err := ioutil.ReadAll(file2)
	if err != nil {
		log.Fatal(err)
	}

	match1, err := regexp.Match(`\r\n`, data1)
	if err != nil {
		log.Fatal(err)
	}
	match2, err := regexp.Match(`\r\n`, data2)
	if err != nil {
		log.Fatal(err)
	}

	text1Platform := "Windows"
	if !match1 {
		text1Platform = "Mac"
	}

	text2Platform := "Windows"
	if !match2 {
		text2Platform = "Mac"
	}

	fmt.Printf("Text1 is using %s line breaks\n", text1Platform)
	fmt.Printf("Text2 is using %s line breaks\n", text2Platform)


}
