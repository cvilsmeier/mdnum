
# mdnum - Markdown Numbered Headings



## Usage


## Usage

~~~
mdnum is a tool for generating numbered headings for markdown

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
    # 0. Chapter 2

    $ mdnum draft.md   # this will overwrite draft.md

    $ cat draft.md
    # 1. Chapter 1
    ## 1.1. Intro
    ## 1.2. Outro
    # 2. Chapter 2
~~~


## Build

You'll need Go installed: https://go.dev

~~~
go install github.com/cvilsmeier/mdnum
~~~

