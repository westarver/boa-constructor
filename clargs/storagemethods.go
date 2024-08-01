package clargs

import (
	"errors"
	"fmt"
	"sort"
)

// exported methods

// for debugging
func (st *ArgList) Print(name string) {
	a := st.Get(name)
	if a != nil {
		fmt.Printf("	name:        %s\n", a.Name())
		fmt.Printf("	param count: %d\n", a.ParamCount())
		fmt.Printf("	param type:  %d\n", a.ParamType())
		fmt.Printf("	short help:  %s\n", a.ShortHelp())
		fmt.Printf("	optional:    %v\n", a.Arg.ParamCount < 0 && a.Arg.ParamCount > -100)
		fmt.Println("	=====================")
	}
}

func (st *ArgList) PrintAll() {
	st.walkAndWorkStr(st.Print)
}

// (*ArgList) Len() returns the number of elements being stored
func (st *ArgList) Len() int {
	return st.len
}

// (*ArgList) Empty() reports if there are 0 elements being stored
func (st *ArgList) Empty() bool {
	return st.len <= 0
}

// Names() returns a slice of strings holding the command names
func (st *ArgList) Names() []string {
	return st.names
}

// (*ArgList) SortedNames() returns a slice of names sorted a-z
func (st *ArgList) SortedNames() []string {
	var names = st.names
	sort.SliceStable(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	return names
}

// (*ArgList) Head() returns the list Head
func (st *ArgList) Head() *CommandLineArg {
	return st.head
}

// (*ArgList) Tail() returns the list Tail
func (st *ArgList) Tail() (*CommandLineArg, bool) {
	if st.last() == nil {
		return nil, false
	}
	return st.last(), true
}

// (*ArgList) Reset()sets current to Head
func (st *ArgList) Reset() {
	if st.head != nil {
		st.current = st.head
	}
}

// (*ArgList) ResetAll() rids the list of all data
func (st *ArgList) ResetAll() {
	st.head = nil
	st.current = nil
	st.names = nil
	st.len = 0
}

// (*ArgList) Get(string) performs a lookup by arg.Name and returns it if found
func (st *ArgList) Get(name string) *CommandLineArg {
	return st.lookup(name)
}

// (*ArgList) Next() returns the next arg and advances current.
// if current has children the head of that list will be next.
func (st *ArgList) Next() (*CommandLineArg, bool) {
	if st.Empty() {
		return nil, false
	}

	if st.current == nil { //bad
		return nil, false
	}

	if st.Len() == 1 {
		return st.Head(), false
	}

	st.current = st.current.next
	return st.current, true
}

// (*ArgList) Previous() returns the previous arg and makes it current
func (st *ArgList) Previous() (*CommandLineArg, bool) {
	if st.current == nil {
		return nil, false
	}
	if st.current == st.head {
		st.current.prev = st.last()
	}
	st.current = st.current.prev
	return st.current, true
}

// (*ArgList) Current() returns the current arg
func (st *ArgList) Current() (*CommandLineArg, bool) {
	if st.current == nil {
		if st.len > 0 {
			st.current = st.head
			return st.current, true
		}
		return nil, false
	}
	return st.current, true
}

// (*ArgList) SetCurrent(string) takes a name and if found makes it the current arg
func (st *ArgList) SetCurrent(name string) bool {
	arg := st.Get(name)
	if arg == nil {
		return false
	}
	st.current = arg
	return true
}

// (*ArgList) Add(arg) appends an arg and makes it current.
// Cannot Add a name that is already in the list
func (st *ArgList) Add(a *CommandLineArg) error {
	if a == nil {
		return errors.New("cannot add nil arg to list")
	}

	if st.Get(a.Name()) != nil {
		return errors.New("cannot duplicate args in a list")
	}

	if st.Empty() {
		a.prev = a
		a.next = a
		st.head = a
		st.current = a
		st.len++
		st.names = append(st.names, a.Name())
		return nil
	}
	last := st.last()
	a.prev = last
	a.next = st.head
	last.next = a

	st.current = a
	st.len++
	st.names = append(st.names, a.Name())
	return nil
}

// // (*ArgList) InsertAfter(string, *CommandLineArg) puts the arg after the named member if found.
// // To insert as Head use the empty string as after
// func (st *ArgList) InsertAfter(after string, a *CommandLineArg) error {
// 	var aft *CommandLineArg
// 	if a == nil {
// 		return errors.New("unable to insert a nil arg")
// 	}

// 	arg := st.Get(a.Name())
// 	if arg != nil {
// 		return errors.New("unable to insert a new arg. use Add instead")
// 	}

// 	if after == "" { // special case of inserting as Head
// 		aft = st.Head() // really before not after
// 		a.next = aft
// 		aft.prev = a
// 		st.head = a
// 		a.prev = nil
// 		st.current = a
// 		st.len++
// 		return nil
// 	}

// 	aft = st.Get(after)
// 	if aft == nil {
// 		return fmt.Errorf("unable to insert after %s, not found", after)
// 	}

// 	tmp := aft.next
// 	aft.next = a
// 	a.prev = aft
// 	a.next = tmp
// 	st.current = a
// 	return nil
// }

// (*ArgList) Delete(string) marks an arg as deleted and if it is current
// advances or retreats as needed. Returns success or failure.
func (st *ArgList) Delete(name string) bool {
	arg := st.Get(name)
	if arg == nil {
		return false
	}

	arg.Arg.IsDeleted = true

	if st.len == 1 {
		st.current = nil
	}
	if st.current == arg {
		if st.current.next != nil {
			st.current = arg.next
			return true
		}
		if st.current.prev != nil {
			st.current = arg.prev
			return true
		}
	}
	arg.Arg.IsDeleted = true
	return true
}

// (*ArgList) UnDelete(string) unmarks an arg as deleted and makes it current
func (st *ArgList) UnDelete(name string) bool {
	arg := st.Get(name)
	if arg == nil {
		return false
	}
	arg.Arg.IsDeleted = false
	st.current = arg
	return true
}

// (*ArgList) IsDeleted(string) returns the deletion status of the named arg
func (st *ArgList) IsDeleted(name string) bool {
	arg := st.Get(name)
	if arg == nil {
		return true
	}
	return arg.Deleted()
}

// (*ArgList) Purge(string) removes a deleted arg from the linked list if found
// this can be dangerous as its children are purged as well
func (st *ArgList) Purge(name string) bool {
	arg := st.Get(name)
	if arg == nil {
		return false
	}
	newcur := st.nextName()
	st.SetCurrent(newcur)
	if st.IsDeleted(name) {
		st.remove(arg)
		return true
	}
	return false
}

// (*ArgList) DeleteAndPurge(string) calls Delete followed by Purge
func (st *ArgList) DeleteAndPurge(name string) bool {
	if st.Delete(name) {
		return st.Purge(name)
	}
	return false
}

// (*ArgList) Update(arg) tries to Add(arg) and if it already exists
// will replace the existing one with the passed arg
func (st *ArgList) Update(a *CommandLineArg) error {
	if a == nil {
		return errors.New("cannot update list with nil arg")
	}

	err := st.Add(a)
	if err != nil { // arg with same name exists
		err = st.updateOldWithNew(a)
		if err != nil {
			return errors.New("cannot update list")
		}
	}
	return nil // nil error
}

// (*ArgList) IsParent checks for children of named arg
func (st *ArgList) IsParent(name string) bool {
	tmp := st.Get(name)
	if tmp == nil {
		return false // not exactly accurate but is childless
	}
	return len(tmp.Children()) == 0 || tmp.Children() == nil
}

// (*ArgList) isChild checks for parent of named arg
func (st *ArgList) IsChild(name string) bool {
	tmp := st.Get(name)
	if tmp == nil {
		return false // not exactly accurate but is parentless
	}
	return tmp.Parent() == ""
}

// (*ArgList) Exists(string) returns true if the name can be found
func (st *ArgList) Exists(name string) bool {
	return st.Get(name) != nil
}

// unexported methods and helpers

func (st *ArgList) nextName() string {
	cur, ok := st.Current()
	if !ok {
		return ""
	}
	name := cur.Name()
	for i, n := range st.names {
		if name == n {
			if i < len(st.names)-1 {
				return st.names[i+1]
			}
			return st.names[0]
		}
	}
	return name
}

func (st *ArgList) remove(arg *CommandLineArg) {
	// remove the name from the names slice
	// the next time the list is written to
	// storage (disk) this one will not be included
	var names = []string{}
	for _, n := range st.names {
		if n == arg.Name() {
			continue
		}
		names = append(names, n)
	}
	st.names = names
	if st.len > 0 {
		st.len--
	}
}

func (st *ArgList) lookup(name string) *CommandLineArg {
	return st.walk(st.Head(), name)
}

func (st *ArgList) walk(c *CommandLineArg, name string) *CommandLineArg {
	for {
		if c == nil {
			return nil
		}

		if st.len == 0 {
			return nil
		}

		if st.len == 1 {
			if name == st.head.Name() {
				return st.head
			}
		}

		if c == st.last() {
			if name == st.last().Name() {
				return st.last()
			}
			return nil
		}

		if c.Name() == name {
			return c
		}
		c = c.next
	}
}

func (st *ArgList) last() *CommandLineArg {
	arg := st.current

	for {
		if arg == nil {
			return nil
		}
		if arg.next == st.head {
			return arg
		}
		arg = arg.next
	}
}

func (st *ArgList) walkAndWorkStr(fn func(string)) {
	arg := st.Head()
	for {
		if arg == nil {
			return
		}
		fn(arg.Name())
		arg = arg.next
	}
}

func (st *ArgList) updateOldWithNew(ne *CommandLineArg) error {
	old := st.Get(ne.Name())

	if old == nil || ne == nil {
		return errors.New("cannot operate on a nil pointer")
	}
	if old.Deleted() {
		return errors.New("cannot update a deleted command/flag")
	}

	old.Arg = ne.Arg
	return nil
}
