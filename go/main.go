package main

import (
	"bufio"
	"bytes"

	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	gzip "github.com/klauspost/pgzip"
)

func main() {

	flag.Int("n", 4, "an integer")
	flag.Parse()

	sourceFile := flag.Arg(0)

	if sourceFile == "" {
		fmt.Println("Dude, you didn't pass in a  file!")
		os.Exit(1)
	}

	fmt.Println("arg 1: ", flag.Arg(0))
	runtime.GOMAXPROCS(runtime.NumCPU())
	file, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}
	r, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(r)
	t := time.Now()

	n := 0
	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)
		s := strings.Split(line, "\n")
		if len(s) < 19 {
			continue
		}
		n++
	}
	d := time.Since(t)
	fmt.Printf("Processed %d entries in %v, %.1f entries/sec.", n, d, float64(n)/(float64(d)/float64(time.Second)))
}

// ReadPackedFile is a function to unpack a tar.gz
func ReadPackedFile(filepath string) {
	if filepath == "" {
		panic("Empty input!")
	}

	processFile(filepath)
}

func processFile(srcFile string) {
	var w bytes.Buffer
	f, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		fmt.Println("GZ NOT")
		os.Exit(1)
	}

	defer gzf.Close()
	fmt.Print(gzf)
	data, err := ioutil.ReadAll(gzf)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)

}
