package client

func (c *Cli) RegisterDeoss(ip string, port int) (string, error) {
	return c.Chain.Register(ip, port)
}
