[![cover.run](https://cover.run/go/github.com/AstroProfundis/tabby.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2FAstroProfundis%2Ftabby)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

# Tabby
A tiny library for super simple Golang tables

This is a fork of [cheynewallace/tabby](https://github.com/cheynewallace/tabby) that uses [juju/ansiterm](https://github.com/juju/ansiterm) instead of `text/tabwriter` as tab writer, so that colored text is supported in the output.

**Get Tabby**
```go
go get github.com/AstroProfundis/tabby
```

**Import Tabby**
```go
import "github.com/AstroProfundis/tabby"
```

Tabby is a tiny (around 70 lines of code) efficient libary for writing extremely simple table based terminal output in Golang.

Many table libraries out there are overly complicated and packed with features you don't need. If you simply want to write clean output to your terminal in table format with minimal effort, Tabby is for you.

Typical examples
* Writing simple tables with heading and tab spaced columns
* Writing log lines to the terminal with evenly spaced columns

## Example With Heading
```go
t := tabby.New()
t.AddHeader("NAME", "TITLE", "DEPARTMENT")
t.AddLine("John Smith", "Developer", "Engineering")
t.Print()
```

**Output**
```
NAME        TITLE      DEPARTMENT
----        -----      ----------
John Smith  Developer  Engineering
```

## Example Without Heading
```go
t := tabby.New()
t.AddLine("Info:", "WEB", "Success 200")
t.AddLine("Info:", "API", "Success 201")
t.AddLine("Error:", "DATABASE", "Connection Established")
t.Print()
```

**Output**
```
Info:   WEB       Success 200
Info:   API       Success 201
Error:  DATABASE  Connection Established
```

## Example With Custom tabWriter
```go
w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
t := tabby.NewCustom(w)
```

## Full Example
```go
package main

import "github.com/AstroProfundis/tabby"

func main() {
	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.Print()
}
```
