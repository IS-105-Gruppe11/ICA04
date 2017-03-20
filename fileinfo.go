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

	f, err  := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	path, _ :=filepath.Abs(filename)


	fmt.Printf("Information about a file: %s \n", path)
	//Størrelse i forskjellig former:
	size := float64(f.Size())
	kibi := size / 1024
	mibi := size / 1048576
	gibi := size / 1073741824
	fmt.Println("Bytes: ", size)
	fmt.Println("Kibibytes: ", kibi)
	fmt.Println("Mibibytes: ", mibi)
	fmt.Println("Gibiibytes: ", gibi)
	// sjekker filens "moder"
	mode := f.Mode()
	fmt.Println("IS a directory: ", mode.IsDir())
	fmt.Println("Is a regular file: ", mode.IsRegular())
	fmt.Println("Filrerights: ", f.Mode())
	fmt.Println("Is append only: ", mode&os.ModeAppend == 1)
	fmt.Println("Is a Unix character device: ", mode&os.ModeCharDevice == 1)
	fmt.Println("Is a Unix block device: ", mode&os.ModeDevice == 1)
	fmt.Println("Is a symbolic link: ", mode&os.ModeSymlink == 1)


}
