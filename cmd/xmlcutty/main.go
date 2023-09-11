// Copyright 2015 by Leipzig University Library, http://ub.uni-leipzig.de
//
//	The Finc Authors, http://finc.info
//	Martin Czygan, <martin.czygan@uni-leipzig.de>
//
// This file is part of some open source application.
//
// Some open source application is free software: you can redistribute
// it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either
// version 3 of the License, or (at your option) any later version.
//
// Some open source application is distributed in the hope that it will
// be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//
// @license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>
//
// xmlcutty is a simple tool for carving out elements from large XML files,
// fast. Since it works in a streaming fashion, it uses almost no memory and
// can process around 1G of XML per minute.
package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/miku/xmlcutty"
	"golang.org/x/net/html/charset"
)

// Version of xmlcutty.
const Version = "0.1.6"

var (
	path    = flag.String("path", "/", "select path")
	root    = flag.String("root", "", "synthetic root element")
	rename  = flag.String("rename", "", "rename wrapper element to this name")
	version = flag.Bool("v", false, "show version")
)

type dummy struct {
	Text []byte `xml:",innerxml"`
}

// lastElement returns the last element of a path like string.
func lastElement(p string) string {
	parts := strings.Split(p, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func main() {
	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	var (
		reader io.Reader
		bw     = bufio.NewWriter(os.Stdout)
	)
	defer bw.Flush()
	if flag.NArg() < 1 {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader = file
	}
	if *path == "/" {
		if *root != "" {
			fmt.Fprintln(bw, "<"+*root+">")
		}
		if _, err := io.Copy(bw, reader); err != nil {
			log.Fatal(err)
		}
		if *root != "" {
			fmt.Fprintln(bw, "</"+*root+">")
		}
		os.Exit(0)
	}
	var (
		stack   = xmlcutty.StringStack{}
		decoder = xml.NewDecoder(reader)
	)
	decoder.Strict = false
	decoder.CharsetReader = charset.NewReaderLabel
	var (
		opener string
		closer string
	)
	switch *rename {
	case "":
		opener = "<" + lastElement(*path) + ">"
		closer = "</" + lastElement(*path) + ">"
	case "\\n":
		opener = "\n"
		closer = ""
	case " ":
		opener = " "
		closer = " "
	default:
		if strings.HasPrefix(*rename, "\\n") {
			opener = strings.Replace(*rename, "\\n", "\n", -1)
			closer = ""
		} else {
			opener = "<" + *rename + ">"
			closer = "</" + *rename + ">"
		}
	}
	if *root != "" {
		fmt.Println("<" + *root + ">")
	}
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch e := t.(type) {
		case xml.StartElement:
			stack.Push(e.Name.Local)
			if *path == stack.String() {
				var d dummy
				if err := decoder.DecodeElement(&d, &e); err != nil {
					log.Fatal(err)
				}
				stack.Pop()
				fmt.Fprint(bw, opener)
				fmt.Fprint(bw, string(d.Text))
				fmt.Fprint(bw, closer)
			}
		case xml.EndElement:
			stack.Pop()
		}
	}
	if *root != "" {
		fmt.Fprintln(bw, "</"+*root+">")
	}
}
