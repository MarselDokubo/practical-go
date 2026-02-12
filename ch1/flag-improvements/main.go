package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type greeterConfig struct {
	numTimes int
	name     string
}

var ErrInvalidPosArgsSpecified = errors.New("More than one positional argument specified.")
var ErrNoName = errors.New("You did not enter your name.")

func parseArgs(w io.Writer, args []string) (greeterConfig, error) {
	c := greeterConfig{}

	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		cmdDesc := `
A greeter application which prints the name you entered a specified number of times.

Usage of %s: <options> [name]`

		fmt.Fprintf(w, cmdDesc, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, ErrInvalidPosArgsSpecified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil
}

func validateArgs(c greeterConfig) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0.")
	}
	return nil
}

func getUserName(r io.Reader, w io.Writer) (string, error) {
	io.WriteString(w, "Welcome. Enter your name: ")
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() && scanner.Err() != nil {
		return "", scanner.Err()
	}
	name := strings.TrimSpace(scanner.Text())

	if len(name) == 0 {
		return "", ErrNoName
	}

	return name, nil
}

func greet(w io.Writer, config greeterConfig) {
	for i := 0; i < config.numTimes; i++ {
		fmt.Fprintf(w, "Howdy %s\n", config.name)
	}
}

func runGreeter(r io.Reader, w io.Writer, c greeterConfig) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getUserName(r, w)
		if err != nil {
			return err
		}
	}
	greet(w, c)
	return nil
}

func main() {
	config, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		if errors.Is(err, ErrInvalidPosArgsSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = runGreeter(os.Stdin, os.Stdin, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
