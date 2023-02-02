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

	fmt.Fscanln(reader, &t)
	for ; t != 0; t-- {
		fmt.Fscanln(reader, &n)

	}
}
