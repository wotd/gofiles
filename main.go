package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		path        string
		name        string
		file_size   string
		file_number int
		dir_number  int
		depth       int
	)

	flag.StringVar(&path, "path", "/tmp", "Path where directories and files will be created")
	flag.StringVar(&name, "name", "ddn", "This string will be part of directories and files names")
	flag.StringVar(&file_size, "file_size", "10M", "Size of files. Can be single also range, like: 10k-102M")

	flag.IntVar(&file_number, "file-number", 10, "Max number of files that will be created")
	flag.IntVar(&dir_number, "dir-number", 3, "Max number of directories that will be created")
	flag.IntVar(&depth, "depth", 5, "Max depth of directories")

	flag.Parse()

	populateDir(path, name, dir_number, depth, file_number, file_size)
}

// Function creates directories and files with specific sizes
// path - where directory structure should be created
// name - string that will be included in directories and files names
// dir_number - max number of directories created in specific location.
//
//	This number is different for every subdirectories
//
// depth - how many levels of subdirectories should be created. This is max number and can be less.
// file_number - Simillar to the dir_number, but for files
// file_size - How big files should be? Can be specific (10M) or range (100k-200M)
//
//	See parseSizes function for details about possible units.
func populateDir(path, name string, dir_number, depth, file_number int, file_size string) {
	if depth > 0 {
		n := rand.Intn(dir_number) + 1
		f_no := rand.Intn(file_number)
		fmt.Println("n: ", n)
		for i := 0; i < n; i++ {
			p := createDir(path, name)
			for c := 0; c < f_no; c++ {
				createFile(p, name, file_size)
			}
			populateDir(p, name, dir_number, depth-1, file_number, file_size)
		}
	}
}

// Function for creating files.
// path - where file should be created
// name - string that will be included in file name
// file_size - How big file should be? Can be specific (10M) or range (100k-200M)
//
//	See parseSizes function for details about possible units.
//
// return pointer to the file
func createFile(path, name, file_size string) *os.File {
	var size int64
	f, err := os.CreateTemp(path, name+"*.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)

	if strings.Contains(file_size, "-") {
		size_ranges := strings.Split(file_size, "-")
		size = rand.Int63n(parseSizes(size_ranges[1])-parseSizes(size_ranges[0])) + parseSizes(size_ranges[0])
	} else {
		size = parseSizes(file_size)
	}
	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}

	return f
}

// Function for creating directory
// path - where firectory should be located
// name - string that will be included in directory name
func createDir(path, name string) string {
	dir, err := os.MkdirTemp(path, name)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// Function to verify size and units for files
// s - string contains size of file.
// How it works: String s is splited to two rune arrays: l and n that contains letters and numbers.
// array n is converted to string and then to integer.
// Array l is converted to string and is checked for possible values:
// k - for kilobytes
// M - for megabytes
// G - for gigabytes
// size of file is multiplied by number of bytes and is retrned as a file size.
func parseSizes(s string) int64 {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}
	size, err := strconv.Atoi(string(n))
	if err != nil {
		log.Fatal(err)
	}
	unit := string(l)
	switch {
	case unit == "k":
		size *= 1_024
	case unit == "M":
		size *= 1_048_576
	case unit == "G":
		size *= 1_073_741_824
	}

	return int64(size)
}

// Function is used to close file at the end of the program.
func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal("error closing file: ", err)
	}
}
