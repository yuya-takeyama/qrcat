package main

import (
	"bytes"
	"flag"
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
)

var reader io.Reader
var buf *bytes.Buffer
var formatStr string
var format Format

func init() {
	var err error

	flag.StringVar(&formatStr, "format", "ascii", "output format (ascii (by default) or png)")
	flag.Parse()

	format = str2format(formatStr)

	buf = new(bytes.Buffer)
	reader, err = argf.From(flag.Args())
	panicIf(err)
}

func main() {
	_, err := buf.ReadFrom(reader)
	panicIf(err)

	code, err := qr.Encode(buf.String(), qr.H)
	panicIf(err)

	if format == ASCII {
		printAsASCII(code)
	} else {
		printAsPNG(code)
	}
}

func printAsPNG(code *qr.Code) {
	os.Stdout.Write(code.PNG())
}

func printAsASCII(code *qr.Code) {
	out := new(bytes.Buffer)

	for i := 0; i < code.Size; i++ {
		for j := 0; j < code.Size; j++ {
			if code.Black(i, j) {
				out.WriteString(ansi.Color("  ", "white:black"))
			} else {
				out.WriteString(ansi.Color("  ", "black:white"))
			}
		}
		out.WriteString("\n")
	}

	out.WriteTo(os.Stdout)
}

func str2format(str string) Format {
	if str == "png" || str == "p" {
		return PNG
	} else {
		return ASCII
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
