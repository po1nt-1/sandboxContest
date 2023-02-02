package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	oddLine := false
	sum := uint64(0)
	t := uint16(0)
	fmt.Fscanln(reader, &t)
	t *= 2
	for ; t > 0; t-- {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		oddLine = !oddLine
		if oddLine {
			continue
		}

		sliceElems := strings.Split(line, " ")

		pMap := make(map[uint64]uint64)
		for _, elem := range sliceElems {
			ui64, _ := strconv.ParseUint(elem, 10, 16)
			pMap[ui64] += 1
		}

		sum = 0
		for key, value := range pMap {
			sum += (value - uint64(value/3)) * key // -2 * key * value
		}
		fmt.Println(sum)
	}
}
