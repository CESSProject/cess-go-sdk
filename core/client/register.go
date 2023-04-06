package client

func (c *Cli) Register(name, multiaddr string, income string, pledge uint64) (string, error) {
	return c.Chain.Register(name, multiaddr, income, pledge)
}
