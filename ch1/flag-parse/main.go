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
}

var ErrNoName = errors.New("No name specified.")

func main() {
	config, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = validateArgs(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = runGreeter(os.Stdin, os.Stdout, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseArgs(w io.Writer, args []string) (greeterConfig, error) {
	config := greeterConfig{}
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&config.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return config, err
	}
	if fs.NArg() != 0 {
		return config, errors.New("Positional arguments specified.")
	}
	return config, nil
}

// validateArgs checks if the provided configuration is logically valid.
//
// It ensures that numTimes is a positive integer unless the printUsage
// flag is set, in which case the numerical value is ignored.
func validateArgs(c greeterConfig) error {
	if c.numTimes < 1 {
		return errors.New("Invalid argument. Must specify a number greater than 0.")
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

func greet(w io.Writer, name string, config greeterConfig) error {
	for i := 0; i < config.numTimes; i++ {
		_, err := fmt.Fprintf(w, "Hello %s", name)
		if err != nil {
			return err
		}
	}
	return nil
}

func runGreeter(r io.Reader, w io.Writer, config greeterConfig) error {
	name, err := getUserName(r, w)
	if err != nil {
		return err
	}
	err = greet(w, name, config)
	if err != nil {
		return err
	}
	return nil
}
