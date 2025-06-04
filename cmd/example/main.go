package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang_injection_test/prefixcalc"
)

//added new coment

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	data, err := io.ReadAll(h.Input)
	if err != nil {
		return err
	}

	result, err := prefixcalc.EvaluatePrefixExpression(strings.TrimSpace(string(data)))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(h.Output, "%.2f\n", result)
	return err
}

func main() {
	expr := flag.String("e", "", "expression string")
	inputFile := flag.String("f", "", "input file path")
	outputFile := flag.String("o", "", "output file path")

	flag.Parse()

	if (*expr == "" && *inputFile == "") || (*expr != "" && *inputFile != "") {
		fmt.Fprintln(os.Stderr, "Must provide exactly one of -e or -f")
		os.Exit(1)
	}

	var input io.Reader
	var output io.Writer = os.Stdout

	if *expr != "" {
		input = strings.NewReader(*expr)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot open input file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot create output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	handler := &ComputeHandler{Input: input, Output: output}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Computation error:", err)
		os.Exit(1)
	}
}
