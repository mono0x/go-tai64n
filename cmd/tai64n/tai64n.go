package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/vektra/tai64n"
)

func processRow(text string) {
	now := tai64n.Now()
	fmt.Println(now.Label(), text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		processRow(text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
