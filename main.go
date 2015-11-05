package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	reader := bufio.NewReaderSize(fp, 1024)
	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		time.Sleep(*duration)
	}
	return nil
}
