// Command phrase-decode decodes phrase from arguments or lines of
// stdin.
//
// Usage:
//
//     phrase-decode PHRASE..
//     phrase-decode <LINES_OF_PHRASES
//
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	mnemonics "gitlab.com/NebulousLabs/entropy-mnemonics"
)

func parse(s string) ([]byte, error) {
	// allow either whitespace or dashes (but just one)

	// for whitespace, be lenient on number of spaces
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return nil, errors.New("empty input")
	}
	if len(fields) == 1 {
		// accept either whitespace or dash-separated input
		fields = strings.Split(s, "-")
	}
	p := mnemonics.Phrase(fields)
	return mnemonics.FromPhrase(p, mnemonics.English)
}

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s PHRASE..\n", prog)
	fmt.Fprintf(flag.CommandLine.Output(), "  %s <LINES_OF_PHRASES\n", prog)
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() > 0 {
		for _, input := range flag.Args() {
			decoded, err := parse(input)
			if err != nil {
				log.Fatalf("decoding input: %q: %v", input, err)
			}
			if _, err := os.Stdout.Write(decoded); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			decoded, err := parse(input)
			if err != nil {
				log.Fatalf("decoding input: %q: %v", input, err)
			}
			if _, err = os.Stdout.Write(decoded); err != nil {
				log.Fatal(err)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("reading standard input: %v", err)
		}
	}
}
