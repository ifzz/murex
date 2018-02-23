package lang

import (
	"errors"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
)

// ShellExitNum is for when running murex in interactive shell mode
var ShellExitNum int

// RunBlockShellNamespace is used for calling code blocks using the shell's
// namespace. eg commands initiated inside the interactive shell
func RunBlockShellNamespace(block []rune, stdin, stdout, stderr stdio.Io) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		proc.ShellProcess, proc.ShellProcess.ScopedVars, proc.ShellProcess.Config,
	)
}

// RunBlockParentNamespace is used for calling code blocks using the parent's
// namespace. eg `source`
func RunBlockParentNamespace(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		proc.ShellProcess, caller.ScopedVars, caller.Config,
	)
}

// RunBlockNewNamespace is for spawning new murex functions. eg `func {}`
func RunBlockNewNamespace(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		caller, caller.ScopedVars.Copy(), caller.Config.Copy(),
	)
}

// RunBlockExistingNamespace is for code blocks as parameters (eg `if {}`,
// `try {}` etc) or inlining code blocks (eg `out @{g *}`)
func RunBlockExistingNamespace(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process) (exitNum int, err error) {
	return processNewBlock(
		block, stdin, stdout, stderr,
		caller, caller.ScopedVars.Copy(), caller.Config,
	)
}

// processNewBlock parses new block and execute the code.
// Inputs are:
//     * the code block ([]rune),
//     * Stdin, stdout and stderr streams; or nil to black hole those data streams,
//     * caller proc.Process to determine scope and any inherited properties.
// Outputs are:
//     * exit number of the last process in the block,
//     * any errors raised during the parse.
//func processNewBlock(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process, conf *config.Config, vars *types.Vars) (exitNum int, err error) {
func processNewBlock(block []rune, stdin, stdout, stderr stdio.Io, caller *proc.Process, vars *types.Vars, conf *config.Config) (exitNum int, err error) {
	debug.Log(string(block))

	//if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
	//	block = block[1 : len(block)-1]
	//}

	container := new(proc.Process)
	container.State = state.MemAllocated
	container.IsBackground = caller.IsBackground
	container.Name = caller.Name
	container.Scope = caller.Scope
	container.Parent = caller
	container.Id = caller.Id
	container.LineNumber = caller.LineNumber
	container.ColNumber = caller.ColNumber
	container.ScopedVars = vars
	container.Config = conf

	//if caller.Name == proc.ShellProcess.Name {
	if caller.Id == proc.ShellProcess.Id {
		container.ExitNum = ShellExitNum
	} else {
		container.RunMode = caller.RunMode
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
	switch container.RunMode {
	case runmode.Normal, runmode.Shell:
		exitNum = runModeNormal(&tree)
	case runmode.Try:
		exitNum = runModeTry(&tree)
	case runmode.TryPipe:
		exitNum = runModeTryPipe(&tree)
	//case runmode.Evil:
	//	panic("Not yet implemented")
	default:
		panic("Unknown run mode")
	}

	// This will just unlock the parent lock. Stdxxx.Close() will still have to be called.
	container.Stdout.UnmakeParent()
	container.Stderr.UnmakeParent()

	//debug.Json("Finished running &tree", tree)
	return
}
