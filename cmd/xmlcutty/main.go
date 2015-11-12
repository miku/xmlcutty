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

const Version = "0.1.1"

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

	fifo := xmlcutty.StringFifo{}
	decoder := xml.NewDecoder(reader)

	var wrapper string
	switch *rename {
	case "":
		wrapper = fmt.Sprintf("<%s>", lastElement(*path))
	case "\\n":
		wrapper = "\n"
	case " ":
		wrapper = " "
	default:
		wrapper = fmt.Sprintf("<%s>", *rename)
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
			fifo.Push(e.Name.Local)
			if *path == fifo.String() {
				var dummy Dummy
				err := decoder.DecodeElement(&dummy, &e)
				if err != nil {
					log.Fatal(err)
				}
				fifo.Pop()
				fmt.Print(wrapper)
				fmt.Print(string(dummy.Text))
				fmt.Print(wrapper)
			}
		case xml.EndElement:
			fifo.Pop()
		}
	}
	if *root != "" {
		fmt.Println("</" + *root + ">")
	}
}
