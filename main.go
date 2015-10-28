package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	file     *string        = flag.String("f", "", "File")
	duration *time.Duration = flag.Duration("d", time.Second, "Dulation")
)

func main() {
	flag.Parse()
	for {
		if err := lineScan(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(*duration)
	}
}

func lineScan() error {
	var fp *os.File
	var err error
	if *file == "" {
		fp = os.Stdin
	} else {
		fp, err = os.Open(*file)
		if err != nil {
			return fmt.Errorf("lineScan os.Open err: %s", err)
		}
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(*duration)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("lineScan scanner.Err err: %s", err)
	}
	return nil
}
