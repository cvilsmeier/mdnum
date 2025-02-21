package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const usage = `mdnum is a tool for generating numbered headings for markdown

Usage
    mdnum FILE

mdnum replaces all ATX ('#') heading numbers appearing in FILE. 
If FILE is -, it reads from stdin and writes to stdout. The
program supports numbering levels up to 6. 

Example:

    $ cat draft.md
    # 0. Chapter 1
    ## 0.0. Intro
    ## 0.0. Outro

    $ mdnum draft.md

    $ cat draft.md
    # 1. Chapter 1
    ## 1.1. Intro
    ## 1.2. Outro

`

func main() {
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("error: missing FILE")
		fmt.Println("")
		fmt.Print(usage)
		os.Exit(2)
	}
	input := readInput(filename)
	output := Convert(input)
	writeOutput(output, filename)
}

func readInput(filename string) string {
	var input []byte
	var err error
	if filename == "-" {
		input, err = io.ReadAll(os.Stdin)
	} else {
		input, err = os.ReadFile(filename)
	}
	if err != nil {
		fatal(err)
	}
	return string(input)
}

func writeOutput(output, filename string) {
	if filename == "-" {
		fmt.Print(output)
		return
	}
	if err := os.WriteFile(filename, []byte(output), 0644); err != nil {
		fatal(err)
	}
}

func fatal(v ...any) {
	fmt.Println(v...)
	os.Exit(1)
}

func Convert(input string) string {
	builder := newBuilder()
	var lineBuf strings.Builder
	for _, r := range input {
		lineBuf.WriteRune(r)
		if r == '\n' {
			builder.addLine(lineBuf.String())
			lineBuf.Reset()
		}
	}
	if lineBuf.Len() > 0 {
		builder.addLine(lineBuf.String())
	}
	return builder.Finish()
}

type Builder struct {
	output  *strings.Builder
	numbers []int
}

func newBuilder() *Builder {
	return &Builder{&strings.Builder{}, []int{0, 0, 0, 0, 0, 0}}
}

func (b *Builder) addLine(line string) {
	b.output.WriteString(b.convertLine(line))
}

func (b *Builder) convertLine(line string) string {
	if !strings.HasPrefix(line, "#") {
		return line
	}
	hashes, numbering, rest, isHeading := b.splitHeading(line)
	if !isHeading {
		return line
	}
	numbering, isNumbering := b.renumber(numbering)
	if !isNumbering {
		return line
	}
	return hashes + " " + numbering + " " + rest
}

func (b *Builder) renumber(numbering string) (string, bool) {
	if numbering == "" {
		return numbering, false
	}
	orig := numbering
	if !strings.HasSuffix(numbering, ".") {
		numbering += "."
	}
	level := -1
	for _, r := range numbering {
		if r == '.' {
			level++
		} else if '0' <= r && r <= '9' {
			// ok
		} else {
			return orig, false
		}
	}
	if level < 0 || 5 < level {
		return orig, false
	}
	numbering = ""
	b.numbers[level]++
	for i := level + 1; i < len(b.numbers); i++ {
		b.numbers[i] = 0
	}
	for i := range level + 1 {
		if b.numbers[i] == 0 {
			b.numbers[i] = 1
		}
		numbering += strconv.Itoa(b.numbers[i]) + "."
	}
	return numbering, true
}

func (b *Builder) splitHeading(line string) (string, string, string, bool) {
	for strings.Contains(line, "  ") {
		line = strings.ReplaceAll(line, "  ", " ")
	}
	toks := strings.SplitN(line, " ", 3)
	if len(toks) != 3 {
		return line, "", "", false
	}
	return toks[0], toks[1], toks[2], true
}

func (b *Builder) Finish() string {
	return b.output.String()
}
