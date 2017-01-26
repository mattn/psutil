package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mattn/psutil"
)

func main() {
	for _, arg := range os.Args[1:] {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		psutil.KillTree(pid, 0)
	}
}
