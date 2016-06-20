// Copyright 2016 Manlio Perillo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// genconf is a simple tool used to generate a configuration file from a
// template (in Go text/template format) using data from a JSON file.
//
// The configuration template is read from stdin and the generated
// configuration file written to stdout.
//
// The template data is read from a JSON file, and stored in a
// map[string]interface{} value.  The data file can be specified using the
// -data flag.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var dataFile = flag.String("data", "", "configuration data file")

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: genconf [flags] [file]")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(2)
	}

	t := parse(input())
	if err := t.Execute(output(), data()); err != nil {
		log.Fatalf("executing template: %v", err)
	}
}

// input returns the File where the configuration template is read from.
func input() *os.File {
	switch {
	case flag.NArg() == 0:
		return os.Stdin
	case flag.Arg(0) == "-":
		return os.Stdin
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("opening template file: %v", err)
	}

	return f
}

// output returns the File where the generated configuration file is written
// to.
func output() *os.File {
	// TODO(mperillo): Add support for the -out flag.
	return os.Stdout
}

// data returns the template data.
func data() map[string]interface{} {
	var data map[string]interface{}

	if *dataFile != "" {
		buf, err := ioutil.ReadFile(*dataFile)
		if err != nil {
			log.Fatalf("reading data file: %v", err)
		}
		if err := json.Unmarshal(buf, &data); err != nil {
			log.Fatalf("unmarshaling data file: %v", err)
		}
	}

	return data
}

// parse creates a new Template and parses the template definitions from the
// File.
func parse(f *os.File) *template.Template {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("reading template: %v", err)
	}
	t, err := template.New(f.Name()).Parse(string(buf))
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	return t
}
