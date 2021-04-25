// Command phrase-encode encodes its standard input as a
// human-friendly phrase.
//
// Usage:
//
//     phrase-encode <FILE
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	mnemonics "gitlab.com/NebulousLabs/entropy-mnemonics"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  %s [OPTIONS] <FILE\n", prog)
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	p, err := mnemonics.ToPhrase(data, mnemonics.English)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Join(p, "-") + "\n"
	if _, err := os.Stdout.WriteString(s); err != nil {
		log.Fatal(err)
	}
}
