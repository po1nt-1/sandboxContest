package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type period struct {
	start uint32
	end   uint32
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	t := uint8(0)
	n := uint16(0)
	no := false

	rawLine := ""
	rawPeriods := make([]string, 6)
	ui32 := uint32(0)
	ui64 := uint64(0)

	fmt.Fscanln(reader, &t)
	for ; t != 0; t-- {
		fmt.Fscanln(reader, &n)

		periods := make([]period, n)
		no = false
	nFor:
		for i := uint16(0); i < n; i++ {
			fmt.Fscanln(reader, &rawLine)

			if no {
				continue
			}

			rawLine = strings.ReplaceAll(rawLine, "-", ":")
			rawPeriods = strings.Split(rawLine, ":")

			for j := 0; j < 6; j++ {
				ui64, _ = strconv.ParseUint(rawPeriods[j], 10, 32)
				ui32 = uint32(ui64)
				switch j {
				case 0:
					if ui32 <= 23 && ui32 >= 0 {
						periods[i].start += ui32 * 60 * 60
					} else {
						no = true
						continue nFor
					}
				case 1:
					if ui32 <= 59 && ui32 >= 0 {
						periods[i].start += ui32 * 60
					} else {
						no = true
						continue nFor
					}
				case 2:
					if ui32 <= 59 && ui32 >= 0 {
						periods[i].start += ui32
					} else {
						no = true
						continue nFor
					}
				case 3:
					if ui32 <= 23 && ui32 >= 0 {
						periods[i].end += ui32 * 60 * 60
					} else {
						no = true
						continue nFor
					}
				case 4:
					if ui32 <= 59 && ui32 >= 0 {
						periods[i].end += ui32 * 60
					} else {
						no = true
						continue nFor
					}
				case 5:
					if ui32 <= 59 && ui32 >= 0 {
						periods[i].end += ui32
					} else {
						no = true
						continue nFor
					}
				}
			}

			if periods[i].start > periods[i].end {
				no = true
			}
		}

		if !no {
			sort.SliceStable(periods, func(i, j int) bool {
				return periods[i].start < periods[j].start
			})

			for i := 1; i < len(periods); i++ {
				if periods[i].start <= periods[i-1].end {
					no = true
					break
				}
			}
		}

		if no {
			fmt.Fprintln(writer, "NO")
		} else {
			fmt.Fprintln(writer, "YES")
		}

	}
}
