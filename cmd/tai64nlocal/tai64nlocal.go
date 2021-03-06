package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vektra/tai64n"
)

func processRow(text string) {
	if !strings.HasPrefix(text, "@") {
		fmt.Printf("%s\r\n", text)
		return
	}
	splited := strings.SplitN(text, " ", 2)
	if len(splited) != 2 {
		fmt.Printf("%s\r\n", text)
		return
	}
	time := tai64n.ParseTAI64NLabel(splited[0])
	if time == nil {
		fmt.Printf("%s\r\n", text)
		return
	}
	fmt.Printf("%s %s\r\n", time.Time().Local().Format("2006-01-02 15:04:05.000000000"), splited[1])
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
