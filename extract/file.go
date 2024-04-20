package extract

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	bz2 "github.com/cosnicolaou/pbzip2"
)

type File struct {
	InPath  string
	OutPath string
	inFile  *os.File
	outFile *os.File
	outCsv  *csv.Writer
	batch   [][]string
}

func (f *File) Validate() *File {
	//1. Check if file exists
	if _, err := os.Stat(f.InPath); errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("file does not exist: %s", f.InPath))
	}
	//2. Check file extension
	if !strings.HasSuffix(f.InPath, "xml.bz2") {
		panic(fmt.Sprintf("file extension is incorrect, expected .xml.bz2: %s",
			f.InPath))
	}
	return f
}

func (f *File) Open() *File {
	//3. Create the input file
	inFile, err := os.OpenFile(f.InPath, 0, 0)
	if err != nil {
		panic(fmt.Sprintf("error while opening the file: %s", f.InPath))
	}
	f.inFile = inFile
	//4. Create the output file
	outFile, err := os.Create(f.OutPath)
	if err != nil {
		panic(fmt.Sprintf("error while creating the file: %s", f.OutPath))
	}
	f.outFile = outFile
	f.outCsv = csv.NewWriter(f.outFile)
	f.batch = make([][]string, 0)
	return f
}

func (f *File) Close() {
	if len(f.batch) >= 0 {
		err := f.outCsv.WriteAll(f.batch)
		if err != nil {
			panic(err)
		}
	}
	err := f.inFile.Close()
	if err != nil {
		panic(err)
	}
	f.outCsv.Flush()
	err = f.outFile.Close()
	if err != nil {
		panic(err)
	}
}

func (f *File) Decompress() *bufio.Scanner {
	pool := bz2.CreateConcurrencyPool(runtime.GOMAXPROCS(-1) - 1)
	// create a reader
	br := bufio.NewReader(f.inFile)
	// create a bzip2.reader, using the reader we just created
	dcr := bz2.NewReader(context.Background(), br,
		bz2.DecompressionOptions(
			bz2.BZConcurrency(runtime.GOMAXPROCS(-1)),
			bz2.BZConcurrencyPool(pool),
		))
	d := bufio.NewReader(dcr)
	buf := make([]byte, 0, 128*1024)
	s := bufio.NewScanner(d)
	s.Buffer(buf, 1024*1024)
	return s
}

func (f *File) WriteRow(row []string) {
	err := f.outCsv.Write(row)
	if err != nil {
		panic(err)
	}
}

func (f *File) BatchWrite(rows []string) {
	f.batch = append(f.batch, rows)
	if len(f.batch) >= 1000 {
		err := f.outCsv.WriteAll(f.batch)
		if err != nil {
			panic(err)
		}
		f.batch = make([][]string, 0)
	}
}
