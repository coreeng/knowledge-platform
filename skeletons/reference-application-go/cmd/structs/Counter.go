package structs

type Counter struct {
	Name    string `json:"name"`
	Counter uint64 `json:"counter"`
}

func (c *Counter) IncrementCounter() {
	c.Counter += 1
}
