package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

// grep --regexp=^root passwd
var pattern string

func init() {
	flag.StringVar(&pattern, "regexp", "", "Specify a pattern")
}

func main() {
	flag.Parse()
	grep(pattern, flag.Args()[0])
}

func grep(pattern, fileName string) {
	exp := regexp.MustCompile(pattern)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if exp.MatchString(line) {
			fmt.Println(line)
		}
	}
}
