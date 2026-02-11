// Greeter (cli) requests user's name and greets them n number of times
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type greeterConfig struct {
	numTimes   int
	printUsage bool
}

var greetDesc = fmt.Sprintf(
	"Usage: %s <integer> [-h|-help] \nA greeter application which prints the name you entered <integer> number of times.\n",
	os.Args[0])

var ErrNoName = errors.New("No name specified.")

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Println(err.Error())
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runGreeter(os.Stdout, os.Stdin, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

}

func printUsage(w io.Writer) {

	io.WriteString(w, greetDesc)
}

// parseArgs processes command-line arguments to configure the greeter.
// It expects a single argument: either a help flag (-h, --help) or
// an integer representing the number of times to greet.
//
// Returns a greeterConfig and an error if the arguments are invalid
// or the input cannot be parsed as an integer.
func parseArgs(args []string) (greeterConfig, error) {
	config := greeterConfig{}
	if len(args) != 1 {
		return config, errors.New("Invalid number of arguments.")
	}
	if args[0] == "-h" || args[0] == "--help" {
		config.printUsage = true
		return config, nil
	}

	numTimes, err := strconv.Atoi(args[0])
	if err != nil {
		return config, err
	}
	config.numTimes = numTimes
	return config, nil
}

// validateArgs checks if the provided configuration is logically valid.
//
// It ensures that numTimes is a positive integer unless the printUsage
// flag is set, in which case the numerical value is ignored.
func validateArgs(c greeterConfig) error {
	if c.numTimes < 1 && !c.printUsage {
		return errors.New("Invalid argument. Must specify a number greater than 0.")
	}
	return nil
}

// getUserName prompts the user for their name via the provided writer
// and reads the input from the provided reader.
//
// It returns the entered name as a string. An error is returned if
// the input fails or if the user provides an empty response.
func getUserName(i io.Reader, o io.Writer) (string, error) {
	var name string
	io.WriteString(o, "Welcome. Enter your name: ")

	scanner := bufio.NewScanner(i)
	if !scanner.Scan() {
		err := scanner.Err()
		// err != EOF
		if err != nil {
			return "", err
		}
	}

	name = strings.TrimSpace(scanner.Text())
	if len(name) == 0 {
		return "", ErrNoName
	}
	return name, nil
}

func greet(o io.Writer, name string, config greeterConfig) {
	greeting := fmt.Sprintf("Howdy %s\n", name)
	for i := 0; i < config.numTimes; i++ {
		io.WriteString(o, greeting)
	}
}

func runGreeter(i io.Reader, w io.Writer, c greeterConfig) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}
	name, err := getUserName(i, w)
	if err != nil {
		return err
	}
	greet(w, name, c)
	return nil
}
