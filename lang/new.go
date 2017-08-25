package lang

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

// ShellExitNum is for when running murex in interactive shell mode
var ShellExitNum int

// ProcessNewBlock parses new block and execute the code.
// Inputs are:
// * the code block ([]rune),
// * Stdin, stdout and stderr streams; or nil to black hole those data streams,
// * caller proc.Process to determine scope and any inherited properties.
// Outputs are:
// * exit number of the last process in the block,
// * any errors raised during the parse.
func ProcessNewBlock(block []rune, stdin, stdout, stderr streams.Io, caller *proc.Process) (exitNum int, err error) {
	container := new(proc.Process)
	container.State = state.MemAllocated
	container.IsBackground = caller.IsBackground
	container.Name = caller.Name
	container.Scope = caller.Scope
	container.Parent = caller
	container.Id = caller.Id
	container.LineNumber = caller.LineNumber
	container.ColNumber = caller.ColNumber
	if caller.Name == proc.ShellProcess.Name {
		container.ExitNum = ShellExitNum
	}

	if stdin != nil {
		container.Stdin = stdin
	} else {
		container.Stdin = streams.NewStdin()
		container.Stdin.SetDataType(types.Null)
		container.Stdin.Close()
	}

	if stdout != nil {
		container.Stdout = stdout
	} else {
		container.Stdout = new(streams.TermOut)
	}
	container.Stdout.MakeParent()

	if stderr != nil {
		container.Stderr = stderr
	} else {
		container.Stderr = new(streams.TermErr)
	}
	container.Stderr.MakeParent()

	tree, pErr := ParseBlock(block)
	if pErr.Code != 0 {
		container.Stderr.Writeln([]byte(pErr.Message))
		//debug.Json("ParseBlock returned:", pErr)
		err = errors.New(pErr.Message)
		return 1, err
	}

	compile(&tree, container)

	// Support for different run modes:
	switch {
	case container.Name == "try":
		exitNum = runHyperSensitive(&tree)
	default:
		exitNum = runNormal(&tree)
	}

	// This will just unlock the parent lock. Stdxxx.Close() will still have to be called.
	container.Stdout.UnmakeParent()
	container.Stderr.UnmakeParent()

	//debug.Json("Finished running &tree", tree)
	return
}
