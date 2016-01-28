package cliwizard

// Context encapsulates a user's input, allowing metadata to be associated with
// it by different engines.
//
// The raw user input must always be available from the Raw commands, whereas
// the contextual methods, are optional. If a context value doesn't exist, it
// will return a nil rather than an empty string.
type Context interface {
	Set(key string, value string)
	Get(key string) *string
	SetRaw(raw string)
	Raw() string
}

type context struct {
	raw  string
	data map[string]*string
}

func newContext() *context {
	return &context{
		data: make(map[string]*string, 0),
	}
}

func (c *context) Set(key string, value string) {
	c.data[key] = &value
}

func (c *context) Get(key string) *string {
	return c.data[key]
}

func (c *context) SetRaw(raw string) {
	c.raw = raw
}

func (c *context) Raw() string {
	return c.raw
}
