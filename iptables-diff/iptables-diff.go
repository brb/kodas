package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func usageAndExit() {
	fmt.Fprintf(os.Stderr, "Usage: %s <before> <after>\n", os.Args[0])
	os.Exit(1)
}

func panicOnErr(ctx string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", ctx, err))
	}
}

func diff(before, after *bufio.Scanner) {
	for before.Scan() && after.Scan() {
		b := strings.Split(before.Text(), " ")
		a := strings.Split(after.Text(), " ")

		if (b[0] == "#" && a[0] == "#") || (strings.HasPrefix(b[0], ":") && strings.HasPrefix(a[0], ":")) || (len(b) == 1 && len(a) == 1) {
			continue

		}
		if len(b) != len(a) {
			panicOnErr("rules do not match", fmt.Errorf("%q vs %q", before.Text(), after.Text()))
		}
		if strings.Join(b[1:], " ") != strings.Join(a[1:], " ") {
			panicOnErr("rules do not match", fmt.Errorf("%q vs %q", before.Text(), after.Text()))
		}

		if b[0][0] == '[' && a[0][0] == '[' {
			bCount, err := strconv.Atoi(strings.Split(strings.Trim(b[0], "[]"), ":")[0])
			panicOnErr("bCount", err)
			aCount, err := strconv.Atoi(strings.Split(strings.Trim(a[0], "[]"), ":")[0])
			panicOnErr("aCount", err)
			d := aCount - bCount
			if d != 0 {
				fmt.Printf("+%d %s\n", d, strings.Join(b[1:], " "))
			}
		}
	}

	if err := before.Err(); err != nil {
		panicOnErr("scan before", err)
	}
	if err := after.Err(); err != nil {
		panicOnErr("scan after", err)
	}
}

func main() {
	if len(os.Args) != 3 {
		usageAndExit()
	}

	before, err := os.Open(os.Args[1])
	panicOnErr("before", err)
	after, err := os.Open(os.Args[2])
	panicOnErr("after", err)

	diff(bufio.NewScanner(before), bufio.NewScanner(after))
}
