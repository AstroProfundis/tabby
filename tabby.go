// This is a fork of github.com/cheynewallace/tabby that uses
// juju/ansiterm.Tabwriter instead of the vanilla one
// This allows us to use colored columns in output

package tabby

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/juju/ansiterm"
)

// Tabby is returned when New() is called.
type Tabby struct {
	writer *ansiterm.TabWriter
}

// New returns a new *ansiterm.TabWriter with default config
func New() *Tabby {
	return &Tabby{
		writer: ansiterm.NewTabWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

// NewCustom returns a new *Tabby with with custom *ansiterm.TabWriter set
func NewCustom(writer *ansiterm.TabWriter) *Tabby {
	return &Tabby{
		writer: writer,
	}
}

// AddLine will write a new table line
func (t *Tabby) AddLine(args ...interface{}) {
	formatString := t.buildFormatString(args)
	fmt.Fprintf(t.writer, formatString, args...)
}

// AddHeader will write a new table line followed by a seperator
func (t *Tabby) AddHeader(args ...interface{}) {
	t.AddLine(args...)
	t.addSeparator(args)
}

// Print will write the table to the terminal
func (t *Tabby) Print() {
	t.writer.Flush()
}

// addSeparator will write a new dash seperator line based on the args length
func (t *Tabby) addSeparator(args []interface{}) {
	var b bytes.Buffer
	for idx, arg := range args {
		length := len(fmt.Sprintf("%v", arg))
		b.WriteString(strings.Repeat("-", length))
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	fmt.Fprintln(t.writer, b.String())
}

// buildFormatString will build up the formatting string used by the *ansiterm.TabWriter
func (t *Tabby) buildFormatString(args []interface{}) string {
	var b bytes.Buffer
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	return b.String()
}
