package readfile

import (
	"bufio"
	"fmt"
	"os"
)

func reader(t string) {
	fmt.Println("Reading file:", t)
	f, err := os.Open(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Initialize a counter
	line := 1
	// Scan through the file line by line
	for scanner.Scan() {
		// Print the line
		fmt.Println(line, scanner.Text())
		line++
		//
	}

}
