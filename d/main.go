package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func printTable(writer *bufio.Writer, table [][]uint8) {
	n := uint8(cap(table))
	m := uint8(cap(table[0]))
	for i := uint8(0); i < n; i++ {
		for j := uint8(0); j < m; j++ {
			fmt.Fprint(writer, table[i][j])
			if j != m-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	t := uint8(0)
	n, m := uint8(0), uint8(0)
	k := uint8(0)

	fmt.Fscanln(reader, &t)
	for ; t != 0; t-- {
		fmt.Fscan(reader, &n, &m)

		table := make([][]uint8, n)
		for i := uint8(0); i < n; i++ {
			table[i] = make([]uint8, m)
			for j := uint8(0); j < m; j++ {
				fmt.Fscan(reader, &table[i][j])
			}
		}

		fmt.Fscan(reader, &k)
		columns := make([]uint8, k)
		for i := uint8(0); i < k; i++ {
			fmt.Fscan(reader, &columns[i])
		}

		for i := uint8(0); i < k; i++ {
			sort.SliceStable(table, func(p, q int) bool {
				return table[p][columns[i]-1] < table[q][columns[i]-1]
			})
		}

		printTable(writer, table)
		fmt.Fprintln(writer)
	}
}
