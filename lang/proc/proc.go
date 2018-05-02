package proc

import (
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/pipes"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/streams/stdio"

	"errors"
	"sync"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
	Stdin              stdio.Io
	Stdout             stdio.Io
	Stderr             stdio.Io
	Parameters         parameters.Parameters
	ExitNum            int
	Name               string
	Id                 int
	Path               string
	IsMethod           bool
	Scope              *Process  `json:"-"`
	Parent             *Process  `json:"-"`
	Previous           *Process  `json:"-"`
	Next               *Process  `json:"-"`
	WaitForTermination chan bool `json:"-"`
	Kill               func()    `json:"-"`
	IsNot              bool
	NamedPipeOut       string
	NamedPipeErr       string
	hasTerminatedM     sync.Mutex
	hasTerminatedV     bool
	State              state.FunctionState
	IsBackground       bool
	LineNumber         int
	ColNumber          int
	RunMode            runmode.RunMode
	Config             *config.Config
	Variables          *Variables
	FidTree            []int
	CreationTime       time.Time
	StartTime          time.Time
	//ScopedVars         *types.Vars
}

var (
	// ShellProcess is the root murex process
	ShellProcess *Process = &Process{}

	// MxFunctions is a table of global murex functions
	MxFunctions MurexFuncs = NewMurexFuncs()

	// GoFunctions is a table of available builtin functions
	GoFunctions map[string]func(*Process) error = make(map[string]func(*Process) error)

	// This will hold all variables
	masterVarTable = newVarTable()

	// InitConf is a table of global config options
	InitConf *config.Config = config.NewConfiguration()

	// GlobalAliases is a table of global aliases
	GlobalAliases Aliases = NewAliases()

	// GlobalPipes is a table of  named pipes
	GlobalPipes pipes.Named = pipes.NewNamed()

	// GlobalFIDs is a table of running murex processes
	GlobalFIDs funcID = *newFuncID()

	// KillForeground is the `kill` function for whichever FID currently has "focus"
	KillForeground func() = func() {}

	// ForegroundProc is the murex FID which currently has "focus"  Em3w
	ForegroundProc *Process = ShellProcess
)

// HasTerminated checks if process has terminated.
// This is a function because terminated state can be subject to race conditions
// so we need a mutex to make the state thread safe.
func (p *Process) HasTerminated() (state bool) {
	p.hasTerminatedM.Lock()
	state = p.hasTerminatedV
	p.hasTerminatedM.Unlock()
	return
}

// SetTerminatedState sets the process terminated state.
// This is a function because terminated state can be subject to race conditions
// so we need a mutex to make the state thread safe.
func (p *Process) SetTerminatedState(state bool) {
	p.hasTerminatedM.Lock()
	p.hasTerminatedV = state
	p.hasTerminatedM.Unlock()
	return
}

// ErrIfNotAMethod returns a standard error message for builtins not run as methods
func (p *Process) ErrIfNotAMethod() (err error) {
	if !p.IsMethod {
		err = errors.New("`" + p.Name + "` expects to be pipelined.")
	}
	return
}

func (p *Process) BranchFID() *Branch {
	process := *p
	process.Name += " (branch)"
	process.Config = process.Config.Copy()
	process.Stdout.Open()
	process.Stderr.Open()

	branch := new(Branch)
	branch.Process = &process

	GlobalFIDs.Register(branch.Process)

	return branch
}

type Branch struct {
	Process *Process
}

func (branch Branch) Close() {
	branch.Process.State = state.Terminating

	branch.Process.Stdout.Close()
	branch.Process.Stderr.Close()

	branch.Process.SetTerminatedState(true)

	GlobalFIDs.Deregister(branch.Process.Id)
	branch.Process.State = state.AwaitingGC
}
