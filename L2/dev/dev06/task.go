package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	nameFile = ""
)

func main() {

	columns := flag.Int("f", 0, "fields")
	delimiter := flag.String("d", "	", "delimiter")
	//separatedOnly := flag.Bool("s", false, "separated")

	flag.Parse()

	file, err := os.Open(nameFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), *delimiter)
		if *columns >= len(str) {
			fmt.Println()
			continue
		}

		fmt.Println(str[*columns-1])
	}
}
