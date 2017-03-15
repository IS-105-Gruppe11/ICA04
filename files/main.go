package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	deepCompare("text1.txt", "text2.txt")
}

func deepCompare(text1, text2 string) bool {
	sf, err := os.Open(text1)
	if err != nil {
		log.Fatal(err)
	}

	df, err := os.Open(text2)
	if err != nil {
		log.Fatal(err)
	}

	sscan := bufio.NewScanner(sf)
	dscan := bufio.NewScanner(df)

	for sscan.Scan() {
		dscan.Scan()
		if !bytes.Equal(sscan.Bytes(), dscan.Bytes()) {
			return true
		}
	}

	return false
}
