package client

func (c *Cli) Exit(role string) (string, error) {
	return c.Chain.Exit(role)
}
