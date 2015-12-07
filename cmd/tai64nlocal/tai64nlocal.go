package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/paulhammond/tai64"
)

func processRow(text string) {
	if !strings.HasPrefix(text, "@") {
		fmt.Println(text)
		return
	}
	splited := strings.SplitN(text, " ", 2)
	if len(splited) != 2 {
		fmt.Println(text)
		return
	}
	time, err := tai64.ParseTai64n(splited[0])
	if err != nil {
		fmt.Println(text)
		return
	}
	fmt.Println(time.Format("2006-01-02 15:04:05.000000000"), splited[1])
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
