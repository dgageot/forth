package forth

import (
	"strings"
)

type command func(r *runner)

type runner struct {
	stack
	commands map[string]command
	in       commands
	out      string
}

// Run runs a sequence of commands.
func Run(input string) string {
	r := runner{
		commands: map[string]command{
			"dump":     dump,
			":":        define,
			"constant": constant,
			".":        print,
			"drop":     drop,
			"dup":      dup,
			"=":        equal,
			"swap":     swap,
			"negate":   negate,
			"abs":      abs,
			"+":        sum,
			"-":        sub,
			"*":        mul,
			"/":        div,
			"mod":      mod,
			"max":      max,
			"min":      min,
		},
		in: commands{
			values: strings.Split(input, " "),
		},
	}

	r.run()

	return r.out
}

func (r *runner) run() {
	for !r.in.empty() {
		command := r.in.pop()
		if command == "bye" {
			return
		}

		r.runCommand(command)
	}
}

func (r *runner) runCommand(v interface{}) {
	if s, ok := v.(string); ok {
		if f, present := r.commands[s]; present {
			f(r)
			return
		}

		r.push(v)
		return
	}

	r.push(v)
}
