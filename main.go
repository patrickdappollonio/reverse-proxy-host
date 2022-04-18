package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
}
