package client

func (c *Cli) Register(name string, income string, pledge uint64) (string, error) {
	return c.Chain.Register(name, c.Multiaddr(), income, pledge)
}
