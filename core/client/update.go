package client

func (c *Cli) Update(name string) (string, error) {
	return c.Chain.Update(name, c.Node.Multiaddr())
}
