package main

import (
	"bufio"
	"fmt"
	"os"
	"readproccesor"
)

type ReadHolder struct {
	Seq1 string
	Seq2 string
}
type PairedRead struct {
	R1 string
	R2 string
}

type Trios struct {
	CellBC string
	Umi    string
	TBC    string
}

func reader(pair PairedRead) map[Trios]int {
	read1 := pair.R1
	read2 := pair.R2
	fmt.Println("Reading Read 1 file:", read1)
	fmt.Println("Reading Read 2 file:", read2)
	// Open the file
	f1, err := os.Open(read1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	f2, err := os.Open(read2)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Close the file and garbage collect the file handle
	defer f1.Close()
	defer f2.Close()
	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)
	// Initialize a counter
	line := 1
	wrong_trio_counter := 0
	// Initialize a map to hold the data
	var trio_holder map[Trios]int
	// Scan through the file line by line
	for scanner1.Scan() && scanner2.Scan() {
		line++
		// Initiate a holder of read sequences and qualities
		var readInfo ReadHolder
		// Add
		if line%4 == 2 {
			read1 := scanner1.Text()
			readInfo.Seq1 = read1
			read2 := scanner2.Text()
			readInfo.Seq2 = read2
		}
		// Pass readinfo to the function to process the read
		trio, err := readproccesor.ProcessRead(readInfo)
		// Handles error
		if err == nil {
			wrong_trio_counter++
		}
		// Check if the trio is already in the map
		if _, ok := trio_holder[trio]; ok {
			trio_holder[trio]++
		} else {
			trio_holder[trio] = 1
		}
	}
	return trio_holder

}

func main() {
	// Get the file name from the command line
	read1_dir := os.Args[1]
	read2_dir := os.Args[2]
	fileName := PairedRead{read1_dir, read2_dir}
	// Call the reader function
	trio_holder := reader(fileName)
	// Print the files
	for key, value := range trio_holder {
		fmt.Println("$v, $v, $v, $v\n", key.CellBC, key.Umi, key.TBC, string(value))
	}
}
