package forth

import "strconv"

type commands struct {
	values []string
}

func (c *commands) empty() bool {
	return len(c.values) == 0
}

func (c *commands) pop() interface{} {
	if c.empty() {
		panic("unable to pop")
	}

	v := c.values[0]
	c.values = c.values[1:]

	if f, err := strconv.ParseFloat(v, 64); err == nil {
		return f
	}
	return v
}
