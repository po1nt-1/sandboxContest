package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	t := uint8(0)
	n := uint16(0)
	answerNo := false
	current := uint16(0)
	last := uint16(0)

	fmt.Fscanln(reader, &t)
	for ; t != 0; t-- {
		fmt.Fscan(reader, &n)

		answerNo = false
		current = uint16(0)
		last = uint16(0)
		hits := make([]uint16, n)
		for i := uint16(0); i < n; i++ {
			fmt.Fscan(reader, &current)
			current--

			if answerNo || hits[current] != 0 && last != current {
				answerNo = true
			} else {
				hits[current]++
				last = current
			}
		}
		if answerNo {
			fmt.Fprintln(writer, "NO")
		} else {
			fmt.Fprintln(writer, "YES")
		}
	}
}
