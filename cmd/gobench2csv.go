package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/tools/benchmark/parse"
)

func main() {
	// Open go bench file.
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Open output csv file.
	outputFile, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("cannot create output csv file", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Start parsing the go bench input.
	set, err := parse.ParseSet(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Write the headers line to the csv file.
	headers := [5]string{"Name", "ns/op", "throughput", "alloc/op", "alloc bytes/op"}
	err = writer.Write(headers[:])
	if err != nil {
		log.Fatal("failed to write csv header", err)
	}

	for _, benchmarks := range set {
		for _, benchmark := range benchmarks {
			fmt.Printf("%s\t%s\n", benchmark.Name, benchmark.String())

			columns := make([]string, 5)
			columns[0] = benchmark.Name
			columns[1] = strconv.FormatFloat(benchmark.NsPerOp, 'f', 2, 64)
			columns[2] = strconv.FormatFloat(benchmark.MBPerS, 'f', 2, 64)
			columns[3] = strconv.FormatUint(benchmark.AllocsPerOp, 10)
			columns[4] = strconv.FormatUint(benchmark.AllocedBytesPerOp, 10)

			err := writer.Write(columns)
			if err != nil {
				log.Fatal("failed to write to csv file", err)
			}
		}
	}

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
