package main

import (
	"testing"

	"github.com/lmorg/murex/lang"
)

// TestMurex tests murex runtime environment can be initialised and and simple
// command line can execute
func TestMurex(t *testing.T) {
	lang.InitEnv()

	block := []rune("a [Mon..Fri]->regexp m/^T/")

	//_, err := lang.RunBlockShellConfigSpace(block, nil, nil, nil)
	_, err := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR).Execute(block)

	if err != nil {
		t.Error(err.Error())
	}
}
