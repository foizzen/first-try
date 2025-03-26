package main

import (
	"fmt"
)

type Command interface {
	Execute() string
}

type CommandWithArgs interface {
	Execute(args ...string) string
}

type CommandWithNum interface {
	Execute() int
}

type CommandRegistry struct {
	com map[string]Command
	comWArgs map[string]CommandWithArgs
	comWNum map[string]CommandWithNum
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		com: make(map[string]Command), 
		comWArgs: make(map[string]CommandWithArgs), 
		comWNum: make(map[string]CommandWithNum),
	}
}

func (c *CommandRegistry) RegisterCommand(name string, cmd interface{}) {
	switch cmd.(type) {
	case Command: 	      c.com[name] = cmd.(Command)
	case CommandWithArgs: c.comWArgs[name] = cmd.(CommandWithArgs)
	case CommandWithNum:  c.comWNum[name] = cmd.(CommandWithNum)
	}
}

func (c *CommandRegistry) RunCommand(name string, txt ...string) interface{} {
	if cmd, exist := c.com[name]; exist {
		return cmd.Execute()
	}
	if cmdWithArgs, exist := c.comWArgs[name]; exist {
		return cmdWithArgs.Execute(txt...)
	}
	if cmdWithNum, exist := c.comWNum[name]; exist {
		return cmdWithNum.Execute()
	}
	return "Unknown command"
}

type SayHello struct {}

func (SayHello) Execute() string {
	return "Hello, world!"
}

type SayBye struct {}

func (SayBye) Execute() string {
	return "Goodbye!"
}

type Toggle struct {
	state bool
}

func (t *Toggle) Execute() string {
	t.state = !t.state
	if t.state { return "ON" }
	return "OFF"
}

type Repeat struct {}

func (r Repeat) Execute(args ...string) string {
	var txt string
	for i:=0; i<3; i++ {
		for _, word := range args {
			txt += word + " "
		}
	}
	return txt
}

type Counter struct {
	count int
}

func (c *Counter) Execute() int {
	c.count++
	return c.count
}

func main() {

	reg := NewCommandRegistry()

	reg.RegisterCommand("hello", SayHello{})
	reg.RegisterCommand("bye", SayBye{})
	
	fmt.Println(reg.RunCommand("hello")) // Hello, world!
	fmt.Println(reg.RunCommand("bye"))   // Goodbye!
	fmt.Println(reg.RunCommand("test"))  // Unknown command

	reg.RegisterCommand("repeat", Repeat{})
	fmt.Println(reg.RunCommand("repeat", "Go!")) // Go! Go! Go!

	reg.RegisterCommand("count", &Counter{})
	fmt.Println(reg.RunCommand("count")) // Counter: 1
	fmt.Println(reg.RunCommand("count")) // Counter: 2
	fmt.Println(reg.RunCommand("count")) // Counter: 3

	reg.RegisterCommand("toggle", &Toggle{state: false})
	fmt.Println(reg.RunCommand("toggle")) // "ON"
	fmt.Println(reg.RunCommand("toggle")) // "OFF"
	fmt.Println(reg.RunCommand("toggle")) // "ON"
	fmt.Println(reg.RunCommand("toggle")) // "OFF"
}

