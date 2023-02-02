package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type developer struct {
	index uint8
	level uint8
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	oddLine := false
	n := uint8(0)
	t := uint8(0)

	_, _ = fmt.Fscanln(reader, &t)
	t *= 2
	for ; t > 0; t-- {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		oddLine = !oddLine
		if oddLine {
			ui64, _ := strconv.ParseUint(line, 10, 8)
			n = uint8(ui64)
			continue
		}

		rawSlice := strings.Split(line, " ")

		developers := make([]developer, n)
		for i, elem := range rawSlice {
			ui64, _ := strconv.ParseUint(elem, 10, 8)
			developers[i].index = uint8(i + 1)
			developers[i].level = uint8(ui64)
		}

		removedIndexes := make(map[uint8]bool, n)
		for _, dev := range developers {
			if _, removed := removedIndexes[dev.index]; removed {
				continue
			}

			neighbours := make([]developer, len(developers))
			copy(neighbours, developers)
			removedIndexes[dev.index] = true

			minDiff := float64(255)
		neighboursRange1:
			for i, neighbour := range neighbours {
				if _, removed := removedIndexes[neighbour.index]; removed {
					continue neighboursRange1
				}

				diff := float64(neighbour.level) - float64(dev.level)
				if diff < 0 {
					diff = -diff
				}
				neighbours[i].level = uint8(diff)
				if diff < minDiff {
					minDiff = diff
				}
			}

		neighboursRange2:
			for _, neighbour := range neighbours {
				if _, removed := removedIndexes[neighbour.index]; removed {
					continue neighboursRange2
				}

				if neighbour.level == uint8(minDiff) {
					if dev.index < neighbour.index {
						fmt.Println(dev.index, neighbour.index)
					} else {
						fmt.Println(neighbour.index, dev.index)
					}
					removedIndexes[neighbour.index] = true
					break
				}
			}
		}
		if t > 1 {
			fmt.Println()
		}
	}
}
