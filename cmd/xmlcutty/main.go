// The game ain't in me no more.
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
)

const Version = "0.1.2"

type Dummy struct {
	Text []byte `xml:",innerxml"`
}

func lastElement(p string) string {
	parts := strings.Split(p, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func main() {
	path := flag.String("path", "/", "select path")
	root := flag.String("root", "", "synthetic root element")
	rename := flag.String("rename", "", "rename wrapper element to this name")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	var reader io.Reader
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
		if _, err := io.Copy(os.Stdout, reader); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	stack := xmlcutty.StringStack{}
	decoder := xml.NewDecoder(reader)

	var opener, closer string
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
		opener = "<" + *rename + ">"
		closer = "</" + *rename + ">"
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
				var dummy Dummy
				err := decoder.DecodeElement(&dummy, &e)
				if err != nil {
					log.Fatal(err)
				}
				stack.Pop()
				fmt.Print(opener)
				fmt.Print(string(dummy.Text))
				fmt.Print(closer)
			}
		case xml.EndElement:
			stack.Pop()
		}
	}
	if *root != "" {
		fmt.Println("</" + *root + ">")
	}
}
