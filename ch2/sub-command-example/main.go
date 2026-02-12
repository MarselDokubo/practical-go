package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var v string
	fs := flag.NewFlagSet("cmd-a", flag.ContinueOnError)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	if err := fs.Parse(os.Args[1:]); err != nil && err != flag.ErrHelp {
		fmt.Println(err)
	}
}

func handleCmdA(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-a", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Executing command A")
	return nil
}

func handleCmdB(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")

	err := fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Executing command B")
	return nil
}
