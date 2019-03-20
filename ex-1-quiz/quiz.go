package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var filename = flag.String("fn", "problems.csv", "the input CSV file with questions and answers")
	var timelimit = flag.Int("t", 30, "time limit in seconds")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Press ENTER to start.....")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	r := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)
	correct, total := 0, 0

outerloop:
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
		ansChannel := make(chan string)

		go func() {
			scanner.Scan()
			ansChannel <- scanner.Text()
		}()

		select {
		case <-timer.C:
			break outerloop
		case answer := <-ansChannel:
			if answer == record[1] {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Wrong!")
			}
		}
	}
	fmt.Printf("\nYou scored %v out of %v\n", correct, total)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
