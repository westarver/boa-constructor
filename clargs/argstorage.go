package clargs

import (
	"github.com/westarver/boa"
)

type CommandLineArg struct {
	Arg        *boa.CmdLineItem
	next, prev *CommandLineArg
} // look in itemmethods.go for impl

type ArgList struct {
	head, current *CommandLineArg // all of the members of this list share a parent
	names         []string
	len           int
} // look in storagemethods for impl

// NewArgList returns a pointer to ArgList.
// ex. MainList := clargs.NewArgList()
func NewArgList() *ArgList {
	return &ArgList{}
}

// NewCommandLineArg constructs and returns a *CommandLineArg
func NewCommandLineArg(arg *boa.CmdLineItem) *CommandLineArg {
	return &CommandLineArg{Arg: arg}
}
