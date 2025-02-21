package main

import (
	"flag"
	"io"
	"log"
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
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	flag.Usage = func() {
		log.Print(usage)
	}
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		log.Print(usage)
		os.Exit(2)
	}
	var input []byte
	var err error
	if filename == "-" {
		input, err = io.ReadAll(os.Stdin)
	} else {
		input, err = os.ReadFile(filename)
	}
	if err != nil {
		log.Fatal(err)
	}
	output := convert(string(input))
	if filename == "-" {
		log.Print(output)
	} else {
		err = os.WriteFile(filename, []byte(output), 0644)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func convert(input string) string {
	builder := NewBuilder()
	for _, line := range splitLines(input) {
		builder.Add(line)
	}
	return builder.Finish()
}

func splitLines(text string) []string {
	var lines []string
	var buf strings.Builder
	for _, r := range text {
		buf.WriteRune(r)
		if r == '\n' {
			lines = append(lines, buf.String())
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		lines = append(lines, buf.String())
	}
	return lines
}

type Builder struct {
	b *strings.Builder
	n []int
}

func NewBuilder() *Builder {
	return &Builder{&strings.Builder{}, []int{0, 0, 0, 0, 0, 0}}
}

func (b *Builder) Add(line string) {
	b.b.WriteString(b.convert(line))
	// b.b.WriteString("\n")
}

func (b *Builder) convert(line string) string {
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
	b.n[level]++
	for i := level + 1; i < len(b.n); i++ {
		b.n[i] = 0
	}
	for i := range level + 1 {
		if b.n[i] == 0 {
			b.n[i] = 1
		}
		numbering += strconv.Itoa(b.n[i]) + "."
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
	return b.b.String()
}
