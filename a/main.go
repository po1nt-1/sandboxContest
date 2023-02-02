package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lineNumbers := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lineElems := strings.Split(line, " ")

		if len(lineElems) == 1 {
			lineNumbers, _ = strconv.Atoi(line)
			continue
		}

		num1, _ := strconv.Atoi(lineElems[0])
		num2, _ := strconv.Atoi(lineElems[1])
		fmt.Println(num1 + num2)

		lineNumbers--
		if lineNumbers <= 0 {
			break
		}
	}

}
