package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	"code.google.com/p/rsc/qr"
	"github.com/mgutz/ansi"
	"github.com/yuya-takeyama/argf"
)

type Format int

const (
	ASCII Format = iota
	PNG
	URI
)

var reader io.Reader
var buf *bytes.Buffer
var formatStr string
var format Format
var displayVersion bool

func init() {
	var err error

	flag.StringVar(&formatStr, "format", "ascii", "output format (ascii|png|uri)")
	flag.BoolVar(&displayVersion, "version", false, "display version")
	flag.Parse()

	format = str2format(formatStr)

	buf = new(bytes.Buffer)
	reader, err = argf.From(flag.Args())
	panicIf(err)
}

func main() {
	if displayVersion {
		fmt.Fprintf(os.Stdout, "qrcat v%s, build %s\n", Version, GitCommit)
		return
	}

	_, err := buf.ReadFrom(reader)
	panicIf(err)

	code, err := qr.Encode(buf.String(), qr.H)
	panicIf(err)

	switch format {
	case ASCII:
		printAsASCII(code)
	case PNG:
		printAsPNG(code)
	case URI:
		printAsURI(code)
	}
}

func printAsPNG(code *qr.Code) {
	os.Stdout.Write(code.PNG())
}

func printAsASCII(code *qr.Code) {
	out := new(bytes.Buffer)

	printBorder(out, code)
	for i := 0; i < code.Size; i++ {
		out.WriteString(ansi.Color("  ", "black:white"))
		for j := 0; j < code.Size; j++ {
			if code.Black(i, j) {
				out.WriteString(ansi.Color("  ", "white:black"))
			} else {
				out.WriteString(ansi.Color("  ", "black:white"))
			}
		}
		out.WriteString(ansi.Color("  ", "black:white"))
		out.WriteString("\n")
	}
	printBorder(out, code)

	out.WriteTo(os.Stdout)
}

func printBorder(out *bytes.Buffer, code *qr.Code) {
	for i := 0; i < code.Size+2; i++ {
		out.WriteString(ansi.Color("  ", "black:white"))
	}
	out.WriteString("\n")
}

func printAsURI(code *qr.Code) {
	out := os.Stdout
	out.WriteString("data:image/png;base64,")
	encoder := base64.NewEncoder(base64.StdEncoding, out)
	encoder.Write(code.PNG())
	encoder.Close()
	out.WriteString("\n")
}

func str2format(str string) Format {
	if str == "png" || str == "p" {
		return PNG
	} else if str == "uri" || str == "url" || str == "u" {
		return URI
	} else {
		return ASCII
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
