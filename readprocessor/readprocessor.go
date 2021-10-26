package readprocessor

import (
	"fmt"
	"io/ioutil"
	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/seq/linear"
	"github.com/biogo/biogo/align"
)

// struct for input reads
type ReadHolder struct {
	Seq1  string
	Seq2  string
}
// Output struct for processed reads
type Trios struct {
	CellBC string
	Umi string
	TBC string
	count int
}

// Slice struct to hold the slice
type Slice struct {
	X int
	Y int
}

func ProcessRead(readInfo *ReadHolder) (Trios, error){
	// get the read sequence
	read1 := readInfo.Seq1
	read2 := readInfo.Seq2
	// get the read quality

	// Initialize the output struct
	var output Trios
	// First get the cell barcodes and 
	// Call fuzzy match function to locate call the barcodes
	// get the cell barcode
	cellBC := read1[0:16]
	// get the UMI
	umi := read[16:28]
	// Check the Q30 score of the read
	output.CellBC = cellBC
	output.Umi = umi
	// Fuzzy match for thr tripBC

	// get the TBC
	tbc := read[12:18]
	// create the output struct
	output := Trios{
		CellBC: cellBC,
		Umi: umi,
		TBC: tbc,
		count: 1,
	}
	return output， err
}

fsa := &linear.Seq{Seq: alphabet.BytesToLetters([]byte("CAATACTGCAGGCCACCTACTCATGCACATAATTGGAAGCGCC"))}
fsa.Alpha = alphabet.DNAgapped
fsb := &linear.Seq{Seq: alphabet.BytesToLetters([]byte("CCGGCCACAACTCGAG"))}
fsb.Alpha = alphabet.DNAgapped
fitted := align.Fitted{
	{0, -1000000, -1000000, -1000000, -1000000},
	{-1000000, 10000, -1000, -1000, -1000},
	{-1000000, -1000, 10000, -1000, -1000},
	{-1000000, -1000, -1000, 10000, -10000},
	{-1000000, -1000,  -1000, -1000, 10000},
}
start := time.Now()
aln, err := fitted.Align(fsa, fsb)
elapsed := time.Since(start)
if err == nil {
	fmt.Printf("%s\n", aln)
	fa := align.Format(fsa, fsb, aln, '-')
	fmt.Printf("%s\n%s\n", fa[0], fa[1])
}