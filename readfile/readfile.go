package main

import (
	"bufio"
	"fmt"
	"os"
	"readproccesor"
)

type ReadHolder struct {
	Seq1  string
	Qual1 string
	Seq2  string
	Qual2 string
}
type PairedRead struct {
	R1 string
	R2 string
}

func reader(pair PairedRead) {
	read1 := pair.R1
	read2 := pair.R2
	fmt.Println("Reading Read 1 file:", read1)
	fmt.Println("Reading Read 2 file:", read2)
	// Open the file
	f1, err := os.Open(read1)
	if err != nil {
		fmt.Println(err)
		return
	}
	f2, err := os.Open(read2)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Close the file and garbage collect the file handle
	defer f1.Close()
	defer f2.Close()
	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)
	// Initialize a counter
	line := 1
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
		if line%4 == 0 {
			readInfo.Qual1 = scanner1.Text()
			readInfo.Qual2 = scanner2.Text()
		}
		// Pass readinfo to the function to process the read
		trio, err := readproccesor.ProcessRead(readInfo)
	}

}

func main() {
	// Get the file name from the command line
	read1_dir := os.Args[1]
	read2_dir := os.Args[2]
	fileName := PairedRead{read1_dir, read2_dir}
	// Call the reader function
	reader(fileName)
}
