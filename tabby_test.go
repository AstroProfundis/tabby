package tabby

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/juju/ansiterm"
)

func Test_buildFormatString(t *testing.T) {
	var b bytes.Buffer
	w := ansiterm.NewTabWriter(&b, 0, 0, 1, '.', 0)
	items := make([]interface{}, 3)
	items[0] = "s1"
	items[1] = "s2"
	items[2] = "s3"

	tabby := NewCustom(w)
	tabby.buildFormatString(items)
	tabby.Print()
	fmtString := b.String()
	if fmtString != fmt.Sprintf("%v.%v.%v\n", items...) {
		t.Errorf("fmtString incorrect, got: %v, want: %v.", fmtString, fmt.Sprintf("%v.%v.%v\n", items...))
	}
}

func Test_New(t *testing.T) {
	tabby := New()
	if tabby.writer == nil {
		t.Errorf("New returning uninitialized writer")
	}
}

func Test_NewCustom(t *testing.T) {
	w := ansiterm.NewTabWriter(os.Stdout, 0, 0, 2, ' ', 0)
	tabby := NewCustom(w)
	if reflect.TypeOf(tabby) != reflect.TypeOf(&Tabby{}) {
		fmt.Println(reflect.TypeOf(tabby))
		t.Errorf("NewCustom incorect type returned")
	}
}

func Test_AddLine(t *testing.T) {
	var b bytes.Buffer
	w := ansiterm.NewTabWriter(&b, 0, 0, 1, '.', 0)
	tabby := NewCustom(w)
	tabby.AddLine("test")
	if b.String() != "test\n" {
		t.Errorf("AddLine not writing to io.Writer")
	}
}

func Test_AddHeader(t *testing.T) {
	var b bytes.Buffer
	w := ansiterm.NewTabWriter(&b, 0, 0, 1, '.', 0)
	tabby := NewCustom(w)
	tabby.AddHeader("test")
	if b.String() != "test\n----\n" {
		t.Errorf("AddHeader not writing to io.Writer")
	}
}
