package clargs

import (
	"github.com/westarver/boa"
	con "github.com/westarver/boa-constructor/constants"
)

func (c CommandLineArg) Name() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.Name
}

func (c CommandLineArg) Alias() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.Alias
}

func (c CommandLineArg) ID() int {
	if c.Arg == nil {
		return -1
	}
	return c.Arg.Id
}

func (c CommandLineArg) Deleted() bool {
	if c.Arg == nil {
		return true
	}
	return c.Arg.IsDeleted
}

func (c CommandLineArg) Kind() con.ArgType {
	if c.Arg == nil {
		return con.ArgInvalid
	}
	if c.Arg.IsFlag {
		return con.ArgFlag
	}
	return con.ArgCommand
}

func (c CommandLineArg) ParamType() boa.ParameterType {
	if c.Arg == nil {
		return -1
	}
	return c.Arg.ParamType
}

func (c CommandLineArg) ParamCount() int {
	if c.Arg == nil {
		return 0
	}
	return c.Arg.ParamCount
}

func (c CommandLineArg) RunCode() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.RunCode
}

func (c CommandLineArg) ShortHelp() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.ShortHelp
}

func (c CommandLineArg) LongHelp() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.LongHelp
}

func (c CommandLineArg) Errors() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.Error()
}

func (c CommandLineArg) DefaultValue() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.DefaultValue
}

func (c CommandLineArg) IsDefault() bool {
	if c.Arg == nil {
		return false
	}
	return c.Arg.IsDefault
}

func (c CommandLineArg) IsParamOpt() bool {
	if c.Arg == nil {
		return false
	}
	return c.ParamCount() < -0 && c.ParamCount() > -100
}

func (c CommandLineArg) IsExclusive() bool {
	if c.Arg == nil {
		return false
	}
	return c.Arg.IsExclusive
}

func (c CommandLineArg) IsRequired() bool {
	if c.Arg == nil {
		return false
	}
	return c.Arg.IsRequired
}

func (c CommandLineArg) Parent() string {
	if c.Arg == nil {
		return ""
	}
	return c.Arg.ParName
}

func (c CommandLineArg) Children() []string {
	if c.Arg == nil {
		return nil
	}
	return c.Arg.ChNames
}
