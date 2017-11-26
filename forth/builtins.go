package forth

import "fmt"

// dump displays the data stack.
func dump(r *runner) {
	r.out += "["

	for i, v := range r.stack.values {
		if i > 0 {
			r.out += ", "
		}
		r.out += fmt.Sprintf("%v", v)
	}

	r.out += "]\n"
}

func print(r *runner) {
	v := r.pop()
	r.out += fmt.Sprintf("%v\n", v)
}

func sum(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	r.push(v1 + v2)
}

func sub(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	r.push(v1 - v2)
}

func mul(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	r.push(v1 * v2)
}

func div(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	r.push(v1 / v2)
}

func mod(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	r.push(float64(int(v1) % int(v2)))
}

func dup(r *runner) {
	v := r.pop()
	r.push(v)
	r.push(v)
}

func swap(r *runner) {
	v1, v2 := r.pop(), r.pop()
	r.push(v1)
	r.push(v2)
}

func equal(r *runner) {
	v1, v2 := r.pop(), r.pop()
	if v1 == v2 {
		r.push("-1")
	} else {
		r.push("0")
	}
}

func negate(r *runner) {
	v := r.popNumber()
	r.push(-v)
}

func abs(r *runner) {
	v := r.popNumber()
	switch {
	case v == -0:
		r.push(0)
	case v <= 0:
		r.push(-v)
	default:
		r.push(v)
	}
}

func max(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	if v1 > v2 {
		r.push(v1)
	} else {
		r.push(v2)
	}
}

func min(r *runner) {
	v1, v2 := r.popNumber(), r.popNumber()
	if v1 < v2 {
		r.push(v1)
	} else {
		r.push(v2)
	}
}

func drop(r *runner) {
	r.pop()
}

func define(r *runner) {
	name := r.in.pop().(string)

	var args []interface{}
	for {
		arg := r.in.pop()
		if arg == ";" {
			break
		}

		args = append(args, arg)
	}

	r.commands[name] = func(r *runner) {
		for _, arg := range args {
			r.runCommand(arg)
		}
	}
}

func constant(r *runner) {
	name := r.in.pop().(string)
	v := r.in.pop()

	r.commands[name] = func(r *runner) {
		r.push(v)
	}
}
