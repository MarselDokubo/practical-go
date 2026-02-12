package main

import (
	"errors"
	"flag"
	"io"
)

var errInvalidSubcommand = errors.New("Invalid Subcommand")
var errMissingSubcommand = errors.New("Missing subcommand")

func main() {

}

// handleCommand parses the command line arguments and calls the appropriate subcommand accordingly
//
//	the first argument is required and is the name of the subcommand [http | grpc | ] to invoke
func handleCommand(w io.Writer, args []string) error {

	fs := flag.NewFlagSet("top-command", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		usage := "Usage: mync [http|grpc] -h\n"
		io.WriteString(w, usage)
		handleHttp(w, []string{"-h"})
		handleGrpc(w, []string{"-h"})
	}
	if err := fs.Parse(args); err != nil && err != flag.ErrHelp {
		return err
	}

	remainingArgs := fs.Args() // ["subcommand", flags, posArgs]
	subCommand := remainingArgs[0]
	subArgs := remainingArgs[1:]
	if fs.NArg() == 0 {
		fs.Usage()
		return errMissingSubcommand
	}

	switch subCommand {
	case "http":
		err := handleHttp(w, subArgs[1:])
		if err != nil {
			return err
		}
	case "grpc":
		err := handleGrpc(w, subArgs[1:])
		if err != nil {
			return err
		}
	default:
		fs.Usage()
		return errInvalidSubcommand
	}
	return nil
}

func handleHttp(w io.Writer, args []string) error {
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		usage := "Usage: mync [http|grpc] -h\n"
		io.WriteString(w, usage)
		handleHttp(w, []string{})
		handleGrpc(w, []string{})
	}
	if err := fs.Parse(args); err != nil && err != flag.ErrHelp {
		return err
	}
	return nil
}

func handleGrpc(w io.Writer, args []string) error {
	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		usage := "Usage: mync [http|grpc] -h\n"
		io.WriteString(w, usage)
		handleHttp(w, []string{})
		handleGrpc(w, []string{})
	}
	if err := fs.Parse(args); err != nil && err != flag.ErrHelp {
		return err
	}
	return nil
}
