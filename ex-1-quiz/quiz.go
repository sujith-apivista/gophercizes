package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var filename = flag.String("fn", "problems.csv", "the input file")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)
	correct, total := 0, 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v ", record[0])
		total++
		scanner.Scan()
		answer := scanner.Text()
		if answer == record[1] {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Wrong!")
		}
	}
	fmt.Printf("You scored %v out of %v\n", correct, total)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
