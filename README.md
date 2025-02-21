
# mdnum - Markdown Numbered Headings


## 1. Download

Download Linux/amd64 or windows/amd64 binary from here:
https://github.com/cvilsmeier/mdnum/releases/latest


## 2. Usage

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


## 3. Build

You'll need Go installed: https://go.dev

~~~
go install github.com/cvilsmeier/mdnum
~~~



## 4. Changelog


### v1.0.0

- first version




## 5. License

~~~
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
~~~
