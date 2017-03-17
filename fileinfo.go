package main

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
)



func main() {
	arg1 := os.Args[1]
	FullFileInfo(arg1)


}


func FullFileInfo(filename string) {

	a, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	path_, _ :=filepath.Abs(filename)


	fmt.Printf("Information about a file: %s \n", path_)
	//Størrelse i forskjellig former:
	size := float64(a.Size())
	kibi := size / 1024
	mibi := size / 1048576
	gibi := size / 1073741824
	fmt.Println("Bytes: ", size)
	fmt.Println("Kibibytes: ", kibi)
	fmt.Println("Mibibytes: ", mibi)
	fmt.Println("Gibiibytes: ", gibi)
	// Switch types mappe eller fil
	switch mode := a.Mode(); {
	case mode.IsRegular():
		fmt.Println("Is a directory")
	case mode.IsDir():
		fmt.Println("Is not a directory")
	case mode&os.ModeAppend != 0:
		fmt.Println("Dette er en snarvei (symbolic link)")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("Dette er en named pipe")
	}
	fmt.Println("Filrettigheter: ", a.Mode())
	mode := a.Mode()
	if mode&os.ModeAppend == 0 {
		fmt.Println("Filen er ikke 'append only'")
	} else {
		fmt.Println("Filen er 'append only'")
	}
	if mode&os.ModeNamedPipe == 0 {
		fmt.Println("Filen er ikke 'name pipe'")
	} else {
		fmt.Println("Filen er 'name pipe'")
	}
	if mode&os.ModeDevice == 0 {
		fmt.Println("Filen er ikke  'device oppgaver'")
	} else {
		fmt.Println("Filen er 'device oppgaver'")
	}
	if mode&os.ModeCharDevice == 0 {
		fmt.Println("Filen er ikke  'Unix character device'")
	} else {
		fmt.Println("Filen er 'Unix character device'")
	}
	if mode&os.ModeSymlink == 0 {
		fmt.Println("Filen er ikke  'symbolic link'")
	} else {
		fmt.Println("Filen er 'symbolic link'")
	}
}
