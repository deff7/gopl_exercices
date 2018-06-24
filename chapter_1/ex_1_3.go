package main

import (
	"fmt"
	"io"
	"strings"
)

func echoConcat(w io.Writer, args []string) {
	var result, sep string
	for _, arg := range args {
		result += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, result)
}

func echoJoin(w io.Writer, args []string) {
	fmt.Fprintln(w, strings.Join(args, " "))
}
