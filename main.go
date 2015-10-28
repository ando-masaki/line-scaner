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
		lineScan()
	}
}

func lineScan() {
	var fp *os.File
	var err error
	if *file == "" {
		fp = os.Stdin
	} else {
		fp, err = os.Open(*file)
		if err != nil {
			log.Fatalf("main os.Open err: %s", err)
		}
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(*duration)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("main scanner.Err err: %s", err)
	}
}
