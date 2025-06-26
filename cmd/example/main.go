package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/zhuravlovO/KPI-APZ-lab2"
)

var (
	expression = flag.String("e", "", "Expression to compute")
	inputFile  = flag.String("f", "", "File with expression to compute")
	outputFile = flag.String("o", "", "File to output the result")
)

func main() {
	flag.Parse()

	if *expression == "" && *inputFile == "" {
		fmt.Fprintln(os.Stderr, "Error: You must provide an expression with -e or an input file with -f")
		os.Exit(1)
	}

	if *expression != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "Error: The -e and -f flags cannot be used at the same time")
		os.Exit(1)
	}

	var reader io.Reader
	if *expression != "" {
		reader = strings.NewReader(*expression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	var writer io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
