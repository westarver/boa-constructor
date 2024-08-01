package args_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/westarver/boa"
	"github.com/westarver/boa-constructor/clargs"
)

func TestStorage_Add_Next(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	list := clargs.NewArgList()

	cmd0 := boa.CmdLineItem{Name: "test0"}
	arg0 := clargs.NewCommandLineArg(&cmd0)
	_ = list.Add(arg0)

	cmd1 := boa.CmdLineItem{Name: "test1"}
	arg1 := clargs.NewCommandLineArg(&cmd1)
	_ = list.Add(arg1)

	cmd2 := boa.CmdLineItem{Name: "test2"}
	arg2 := clargs.NewCommandLineArg(&cmd2)
	_ = list.Add(arg2)

	cmd3 := boa.CmdLineItem{Name: "test3"}
	arg3 := clargs.NewCommandLineArg(&cmd3)
	_ = list.Add(arg3)

	//---------------------------

	_ = list.SetCurrent(list.Head().Name())
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "no dice")
	}
	check("", "test0", buf, t)

	if arg21, ok := list.Next(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "next failed")
	}
	check("", "test1", buf, t)

	if arg21, ok := list.Next(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "next failed")
	}
	check("", "test2", buf, t)

	if arg21, ok := list.Next(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "next failed")
	}
	check("", "test3", buf, t)

	if arg21, ok := list.Next(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "next failed")
	}
	check("", "test0", buf, t)
	// all above pass

}

func TestStorage_Add_Previous(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	list := clargs.NewArgList()

	cmd0 := boa.CmdLineItem{Name: "test0"}
	arg0 := clargs.NewCommandLineArg(&cmd0)
	_ = list.Add(arg0)

	cmd1 := boa.CmdLineItem{Name: "test1"}
	arg1 := clargs.NewCommandLineArg(&cmd1)
	_ = list.Add(arg1)

	cmd2 := boa.CmdLineItem{Name: "test2"}
	arg2 := clargs.NewCommandLineArg(&cmd2)
	_ = list.Add(arg2)

	cmd3 := boa.CmdLineItem{Name: "test3"}
	arg3 := clargs.NewCommandLineArg(&cmd3)
	_ = list.Add(arg3)

	//---------------------------

	_ = list.SetCurrent(list.Head().Name())
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "no dice")
	}
	check("0", "test0", buf, t)

	if arg21, ok := list.Previous(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("3", "test3", buf, t)

	if arg21, ok := list.Previous(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("2", "test2", buf, t)

	if arg21, ok := list.Previous(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("1", "test1", buf, t)

	if arg21, ok := list.Previous(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("0", "test0", buf, t)
	// all above pass
}

func TestStorage_Current_SetCurrent(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	list := clargs.NewArgList()

	cmd0 := boa.CmdLineItem{Name: "test0"}
	arg0 := clargs.NewCommandLineArg(&cmd0)
	_ = list.Add(arg0)

	cmd1 := boa.CmdLineItem{Name: "test1"}
	arg1 := clargs.NewCommandLineArg(&cmd1)
	_ = list.Add(arg1)

	cmd2 := boa.CmdLineItem{Name: "test2"}
	arg2 := clargs.NewCommandLineArg(&cmd2)
	_ = list.Add(arg2)

	cmd3 := boa.CmdLineItem{Name: "test3"}
	arg3 := clargs.NewCommandLineArg(&cmd3)
	_ = list.Add(arg3)

	//---------------------------

	_ = list.SetCurrent(list.Head().Name())
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "no dice")
	}
	check("0", "test0", buf, t)

	_ = list.SetCurrent("test1")
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("1", "test1", buf, t)

	_ = list.SetCurrent("test2")
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("2", "test2", buf, t)

	_ = list.SetCurrent("test3")
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("3", "test3", buf, t)

	_ = list.SetCurrent("test0")
	if arg21, ok := list.Current(); ok {
		print(buf, arg21.Name())
	} else {
		print(buf, "failed")
	}
	check("0", "test0", buf, t)
	// all above pass
}

// func TestStorage_InsertAfter(t *testing.T) {
// 	t.Parallel()
// 	buf := &bytes.Buffer{}
// 	list := clargs.NewArgList(clargs.DummyList)

// 	cmd0 := boa.CmdLineItem{Name: "test0"}
// 	arg0 := clargs.NewCommandLineArg(&cmd0)
// 	_ = list.Add(arg0)

// 	cmd1 := boa.CmdLineItem{Name: "test1"}
// 	arg1 := clargs.NewCommandLineArg(&cmd1)
// 	_ = list.Add(arg1)

// 	cmd2 := boa.CmdLineItem{Name: "test2"}
// 	arg2 := clargs.NewCommandLineArg(&cmd2)
// 	_ = list.Add(arg2)

// 	cmd3 := boa.CmdLineItem{Name: "test3"}
// 	arg3 := clargs.NewCommandLineArg(&cmd3)
// 	_ = list.Add(arg3)

// 	//---------------------------

// 	cmd25 := boa.CmdLineItem{Name: "test2.5"}
// 	arg25 := clargs.NewCommandLineArg(&cmd25)
// 	_ = list.InsertAfter("test2", arg25)
// 	if arg99, ok := list.Current(); ok {
// 		print(buf, arg99.Name())
// 	} else {
// 		print(buf, "failed")
// 	}
// 	check("2.5", "test2.5", buf, t)

// 	cmd35 := boa.CmdLineItem{Name: "test3.5"}
// 	arg35 := clargs.NewCommandLineArg(&cmd35)
// 	err := list.InsertAfter("bogus", arg35)
// 	if err == nil {
// 		print(buf, "test3.5")
// 	} else {
// 		print(buf, err.Error())
// 	}
// 	check("3.5", "unable to insert after bogus, not found", buf, t)

// 	cmd05 := boa.CmdLineItem{Name: "test0.5"}
// 	arg05 := clargs.NewCommandLineArg(&cmd05)
// 	_ = list.InsertAfter("", arg05)
// 	if arg99, ok := list.Current(); ok {
// 		print(buf, arg99.Name())
// 	} else {
// 		print(buf, "failed")
// 	}
// 	check("0.5", "test0.5", buf, t)

// 	// all above pass
// }

func TestStorage_Delete_UnDelete_Purge(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	list := clargs.NewArgList()

	cmd0 := boa.CmdLineItem{Name: "test0"}
	arg0 := clargs.NewCommandLineArg(&cmd0)
	_ = list.Add(arg0)

	cmd1 := boa.CmdLineItem{Name: "test1"}
	arg1 := clargs.NewCommandLineArg(&cmd1)
	_ = list.Add(arg1)

	cmd2 := boa.CmdLineItem{Name: "test2"}
	arg2 := clargs.NewCommandLineArg(&cmd2)
	_ = list.Add(arg2)

	cmd3 := boa.CmdLineItem{Name: "test3"}
	arg3 := clargs.NewCommandLineArg(&cmd3)
	_ = list.Add(arg3)

	//---------------------------
	b := list.Delete("test2")
	print(buf, fmt.Sprintf("%v", b))
	check("del", "true", buf, t)

	b = list.IsDeleted("test2")
	print(buf, fmt.Sprintf("%v", b))
	check("is del", "true", buf, t)

	b = list.UnDelete("test2")
	print(buf, fmt.Sprintf("%v", b))
	check("undel", "true", buf, t)

	b = list.IsDeleted("test2")
	print(buf, fmt.Sprintf("%v", b))
	check("is del 2", "false", buf, t)

	// all above pass
}

func TestStorage_Purge(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	list := clargs.NewArgList()

	cmd0 := boa.CmdLineItem{Name: "test0"}
	arg0 := clargs.NewCommandLineArg(&cmd0)
	_ = list.Add(arg0)

	cmd1 := boa.CmdLineItem{Name: "test1"}
	arg1 := clargs.NewCommandLineArg(&cmd1)
	_ = list.Add(arg1)

	cmd2 := boa.CmdLineItem{Name: "test2"}
	arg2 := clargs.NewCommandLineArg(&cmd2)
	_ = list.Add(arg2)

	cmd3 := boa.CmdLineItem{Name: "test3"}
	arg3 := clargs.NewCommandLineArg(&cmd3)
	_ = list.Add(arg3)

	//---------------------------
	b := list.Delete("test2")
	print(buf, fmt.Sprintf("%v", b))
	check("del", "true", buf, t)

	b = list.Purge("test2")
	print(buf, fmt.Sprintf("%v", b))
	check("purge 2", "true", buf, t)

	arg99 := list.Get("test2")
	if arg99 == nil {
		print(buf, "it worked")
	} else {
		print(buf, arg99.Name())
	}
	check("get 2", "it worked", buf, t)
	// all above fail
}

func check(id, w string, buf *bytes.Buffer, t *testing.T) {
	want := w
	got := buf.String()

	if want != got {
		t.Errorf(" %s failed", id)
	}
}

func print(buf *bytes.Buffer, got string) {
	buf.Reset()
	fmt.Fprintf(buf, "%s", got)
}
