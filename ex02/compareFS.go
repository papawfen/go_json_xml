package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func compareFS(path1 string, path2 string) {
	set := make(map[string]bool)

	file1, err := os.Open(path1)
	if err != nil {
		printError("Error: cant open file 1")
	}
	
	scanner1 := bufio.NewScanner(file1)
	for scanner1.Scan() {
		line := scanner1.Text()
		set[line] = true
	}
	file1.Close()

	file2, err := os.Open(path2)
	if err != nil {
		printError("Error: cant open file 2")
	}

	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := scanner2.Text()
		if !set[line] {
			fmt.Println("ADDED " + line)
		}
	}
	
	for k := range set {
		delete(set, k)
	}
	
	file2.Seek(0, io.SeekStart)
	scanner2 = bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := scanner2.Text()
		set[line] = true
	}
	file2.Close()

	file1, err = os.Open(path1)
	if err != nil {
		printError("Error: cant open file 1")
	}
	scanner1 = bufio.NewScanner(file1)
	for scanner1.Scan() {
		line := scanner1.Text()
		if !set[line] {
			fmt.Println("REMOVED " + line)
		}
	}
}


func main() {
	if len(os.Args) == 1 {
		printError("./compareFS --old snapshot1.txt --new snapshot2.txt")
	} else {
		args := os.Args[1:]
		if args[0] != "--old" {
			printError("./compareFS --old snapshot1.txt --new snapshot2.txt")
		} else if args[2] != "--new" {
			printError("./compareFS --old snapshot1.txt --new snapshot2.txt")
		}
		if len(args) < 4 {
			printError("./compareFS --old snapshot1.txt --new snapshot2.txt")
		}
		if strings.HasSuffix(args[1], ".txt") && strings.HasSuffix(args[3], ".txt") {
			compareFS(args[1], args[3])
		} else {
			printError("./compareFS --old snapshot1.txt --new snapshot2.txt")
		}
	}
}